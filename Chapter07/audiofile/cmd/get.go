package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/marianina8/audiofile/utils"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get audio metadata",
	Long:  `Get audio metadata by audiofile id.  Metadata includes available tags and transcript.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		b, err := getAudioByID(cmd)
		if err != nil {
			return err
		}
		fmt.Println(string(b))
		return nil
	},
}

func init() {
	getCmd.Flags().String("id", "", "audiofile id")
	rootCmd.AddCommand(getCmd)
}

func getAudioByID(cmd *cobra.Command) ([]byte, error) {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}
	id, err := cmd.Flags().GetString("id")
	if err != nil {
		fmt.Printf("error retrieving id: %s\n", err.Error())
		return nil, err
	}
	params := "id=" + url.QueryEscape(id)
	path := fmt.Sprintf("http://localhost/request?%s", params)
	payload := &bytes.Buffer{}

	req, err := http.NewRequest(http.MethodGet, path, payload)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = utils.CheckResponse(resp)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}
