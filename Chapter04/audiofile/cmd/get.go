package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get audio metadata",
	Long:  `Get audio metadata by audiofile id.  Metadata includes available tags and transcript.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := &http.Client{}
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			fmt.Printf("error retrieving id: %s\n", err.Error())
			return err
		}
		params := "id=" + url.QueryEscape(id)
		path := fmt.Sprintf("http://localhost/request?%s", params)
		payload := &bytes.Buffer{}

		req, err := http.NewRequest(http.MethodGet, path, payload)
		if err != nil {
			return err
		}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		b, err := io.ReadAll(resp.Body)
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
