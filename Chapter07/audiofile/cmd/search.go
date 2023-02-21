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

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Command to search for audiofiles by string",
	Long:  `Command to search for audiofiles by search string within the metadata file.  Search string is not case sensitive`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := &http.Client{
			Timeout: 15 * time.Second,
		}
		value, err := cmd.Flags().GetString("value")
		if err != nil {
			fmt.Printf("error retrieving value: %s\n", err.Error())
			return err
		}
		params := "searchFor=" + url.QueryEscape(value)
		path := fmt.Sprintf("http://localhost/search?%s", params)
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
		err = utils.CheckResponse(resp)
		if err != nil {
			return err
		}
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		fmt.Println(string(b))
		return nil
	},
}

func init() {
	searchCmd.Flags().String("value", "", "string to search for in metadata")
	rootCmd.AddCommand(searchCmd)
}
