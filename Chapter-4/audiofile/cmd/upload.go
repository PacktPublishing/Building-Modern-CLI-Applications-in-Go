package cmd

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload an audio file",
	Long: `Upload an audio file by passing in the --filename or -f flag followed by the 
filepath of the audiofile.`,
	SuggestFor: []string{"add"},
	RunE: func(cmd *cobra.Command, args []string) error {
		client := &http.Client{}
		filename, err := cmd.Flags().GetString("filename")
		if err != nil {
			fmt.Printf("error retrieving filename: %s\n", err.Error())
			return err
		}
		fmt.Println("Uploading", filename, "...")
		url := "http://localhost/upload"
		payload := &bytes.Buffer{}
		multipartWriter := multipart.NewWriter(payload)
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		partWriter, err := multipartWriter.CreateFormFile("file", filepath.Base(filename))
		if err != nil {
			return err
		}

		_, err = io.Copy(partWriter, file)
		if err != nil {
			return err
		}

		err = multipartWriter.Close()
		if err != nil {
			return err
		}
		req, err := http.NewRequest(http.MethodPost, url, payload)
		if err != nil {
			return err
		}

		req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
		res, err := client.Do(req)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		fmt.Println("Audiofile ID: ", string(body))
		return nil
	},
}

func init() {
	uploadCmd.Flags().StringP("filename", "f", "", "Filepath of filename to be uploaded")
	rootCmd.AddCommand(uploadCmd)
}
