package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/marianina8/audiofile/models"
	"github.com/marianina8/audiofile/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all audio files",
	Long: `List audio file metadata in JSON format.  Data includes id, tags, 
and transcript if available.`,
	Example: `audiofile list`,
	RunE: func(cmd *cobra.Command, args []string) error {
		verbose, _ := cmd.Flags().GetBool("verbose")
		b, err := callList(verbose)
		if err != nil {
			return err
		}
		plainFormat, _ := cmd.Flags().GetBool("plain")
		if plainFormat {
			var audios models.AudioList
			json.Unmarshal(b, &audios)
			fmt.Fprintf(cmd.OutOrStdout(), audios.Plain())
			return nil
		}
		jsonFormat, _ := cmd.Flags().GetBool("json")
		formatedBytes, err := utils.Print(b, jsonFormat)
		if err != nil {
			fmt.Fprintf(cmd.OutOrStdout(), string(formatedBytes))
		}
		return nil
	},
}

func callList(verbose bool) ([]byte, error) {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	path := fmt.Sprintf("http://%s:%d/list", viper.Get("cli.hostname"), viper.GetInt("cli.port"))
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

func init() {
	listCmd.Flags().Bool("json", false, "return json format")
	listCmd.Flags().Bool("plain", false, "return plain format")
	rootCmd.AddCommand(listCmd)
}
