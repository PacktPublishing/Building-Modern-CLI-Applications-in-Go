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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete audiofile by id",
	Long:  `Delete audiofile by id. This command removes the entire folder containing all stored metadata`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := &http.Client{
			Timeout: 15 * time.Second,
		}
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			fmt.Printf("error retrieving id: %s\n", err.Error())
			return err
		}
		params := "id=" + url.QueryEscape(id)
		path := fmt.Sprintf("http://localhost/delete?%s", params)
		payload := &bytes.Buffer{}

		req, err := http.NewRequest(http.MethodDelete, path, payload)
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
	deleteCmd.Flags().String("id", "", "audiofile id")
	rootCmd.AddCommand(deleteCmd)
}
