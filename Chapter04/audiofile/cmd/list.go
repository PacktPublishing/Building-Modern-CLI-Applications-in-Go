package cmd

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all audio files",
	Long: `List audio file metadata in JSON format.  Data includes id, tags, 
and transcript if available.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := &http.Client{}

		path := "http://localhost/list"
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
	rootCmd.AddCommand(listCmd)
}
