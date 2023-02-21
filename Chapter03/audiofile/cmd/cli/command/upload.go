package command

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/marianina8/audiofile/internal/interfaces"
)

func NewUploadCommand(client interfaces.Client) *UploadCommand {
	gc := &UploadCommand{
		fs:     flag.NewFlagSet("upload", flag.ContinueOnError),
		client: client,
	}

	gc.fs.StringVar(&gc.filename, "filename", "", "full path of filename to be uploaded")

	return gc
}

type UploadCommand struct {
	fs       *flag.FlagSet
	client   interfaces.Client
	filename string
}

func (cmd *UploadCommand) Name() string {
	return cmd.fs.Name()
}

func (cmd *UploadCommand) ParseFlags(flags []string) error {
	if len(flags) == 0 {
		fmt.Println("usage: ./audiofile-cli upload -filename <filename>")
		return fmt.Errorf("missing flags")
	}
	return cmd.fs.Parse(flags)
}

func (cmd *UploadCommand) Run() error {
	if cmd.filename == "" {
		return fmt.Errorf("missing filename")
	}
	fmt.Println("Uploading", cmd.filename, "...")
	url := "http://localhost/upload"
	payload := &bytes.Buffer{}
	multipartWriter := multipart.NewWriter(payload)
	file, err := os.Open(cmd.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	partWriter, err := multipartWriter.CreateFormFile("file", filepath.Base(cmd.filename))
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

	client := cmd.client
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
}
