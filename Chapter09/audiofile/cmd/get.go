package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/marianina8/audiofile/models"
	"github.com/marianina8/audiofile/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get audio metadata",
	Long:    `Get audio metadata by audiofile id.  Metadata includes available tags and transcript.`,
	Example: `audiofile get --id 45705eba-9342-4952-8cd4-baa2acc25188`,
	RunE: func(cmd *cobra.Command, args []string) error {
		verbose, _ := cmd.Flags().GetBool("verbose")
		b, err := getAudioByID(cmd, verbose)
		if err != nil {
			return err
		}
		plainFormat, _ := cmd.Flags().GetBool("plain")
		if plainFormat {
			var audio models.Audio
			json.Unmarshal(b, &audio)
			fmt.Fprintf(cmd.OutOrStdout(), audio.Plain())
			return nil
		}
		jsonFormat, _ := cmd.Flags().GetBool("json")
		formattedBytes, err := utils.Print(b, jsonFormat)
		if err != nil {
			fmt.Fprintf(cmd.OutOrStdout(), string(formattedBytes))
		}
		return nil
	},
}

func init() {
	getCmd.Flags().String("id", "", "audiofile id")
	getCmd.Flags().Bool("json", false, "return json format")
	getCmd.Flags().Bool("plain", false, "return plain format")
	rootCmd.AddCommand(getCmd)
}

func getAudioByID(cmd *cobra.Command, verbose bool) ([]byte, error) {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	var err error
	id, _ := cmd.Flags().GetString("id")
	if id == "" {
		id, err = utils.AskForID()
		if err != nil {
			return nil, utils.Error("\n  %v\n  try again and enter an id", err, verbose)
		}
	}
	params := "id=" + url.QueryEscape(id)
	path := fmt.Sprintf("http://%s:%d/request?%s", viper.Get("cli.hostname"), viper.GetInt("cli.port"), params)
	payload := &bytes.Buffer{}
	req, err := http.NewRequest(http.MethodGet, path, payload)
	if err != nil {
		return nil, utils.Error("\n  %v\n  check configuration to ensure properly configured hostname and port", err, verbose)
	}
	utils.LogRequest(verbose, http.MethodGet, path, payload.String())
	resp, err := client.Do(req)
	if err != nil {
		return nil, utils.Error("\n  %v\n  check configuration to ensure properly configured hostname and port\n  or check that api is running", err, verbose)
	}
	defer resp.Body.Close()
	err = utils.CheckResponse(resp)
	if err != nil {
		return nil, utils.Error("\n  checking response: %v", err, verbose)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, utils.Error("\n  reading response: %v\n  ", err, verbose)
	}
	utils.LogHTTPResponse(verbose, resp, b)
	return b, nil
}
