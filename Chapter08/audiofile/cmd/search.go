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

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Command to search for audiofiles by string",
	Long:  `Command to search for audiofiles by search string within the metadata file.  Search string is not case sensitive`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := &http.Client{
			Timeout: 15 * time.Second,
		}
		var err error
		value, _ := cmd.Flags().GetString("value")
		if value == "" {
			value, err = utils.AskForValue()
			if err != nil {
				return err
			}
		}
		params := "searchFor=" + url.QueryEscape(value)
		path := fmt.Sprintf("http://%s:%d/search?%s", viper.Get("cli.hostname"), viper.GetInt("cli.port"), params)
		payload := &bytes.Buffer{}

		req, err := http.NewRequest(http.MethodGet, path, payload)
		if err != nil {
			return err
		}
		fmt.Printf("Sending request: %s %s %s...\n", http.MethodGet, path, payload)
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
		plainFormat, _ := cmd.Flags().GetBool("plain")
		if plainFormat {
			var audios models.AudioList
			json.Unmarshal(b, &audios)
			fmt.Fprintf(cmd.OutOrStdout(), audios.Plain())
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
	searchCmd.Flags().String("value", "", "string to search for in metadata")
	searchCmd.Flags().Bool("json", false, "return json format")
	searchCmd.Flags().Bool("plain", false, "return plain format")
	rootCmd.AddCommand(searchCmd)
}
