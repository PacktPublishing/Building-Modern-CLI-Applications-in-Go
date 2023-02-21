package command

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"

	"github.com/marianina8/audiofile/internal/interfaces"
)

func NewListCommand(client interfaces.Client) *ListCommand {
	gc := &ListCommand{
		fs:     flag.NewFlagSet("list", flag.ContinueOnError),
		client: client,
	}
	return gc
}

type ListCommand struct {
	fs     *flag.FlagSet
	client interfaces.Client
}

func (cmd *ListCommand) Name() string {
	return cmd.fs.Name()
}

func (cmd *ListCommand) ParseFlags(flags []string) error {
	return cmd.fs.Parse(flags)
}

func (cmd *ListCommand) Run() error {
	path := "http://localhost/list"
	payload := &bytes.Buffer{}
	client := cmd.client

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
}
