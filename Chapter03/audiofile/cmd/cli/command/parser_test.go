package command

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/marianina8/audiofile/internal/interfaces"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	if strings.Contains(req.URL.String(), "/upload") {
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader("123")),
		}, nil
	}
	if strings.Contains(req.URL.String(), "/request") {
		value, ok := req.URL.Query()["id"]
		if !ok || len(value[0]) < 1 {
			return &http.Response{
				StatusCode: 500,
				Body:       io.NopCloser(strings.NewReader("url param 'id' is missing")),
			}, fmt.Errorf("url param 'id' is missing")
		}
		if string(value[0]) != "123" {
			return &http.Response{
				StatusCode: 500,
				Body:       io.NopCloser(strings.NewReader("audiofile id does not exist")),
			}, fmt.Errorf("audiofile id does not exist")
		}
		file, err := os.ReadFile("testdata/audio.json")
		if err != nil {
			return nil, err
		}
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(strings.NewReader(string(file))),
		}, nil
	}
	return nil, nil
}

func TestParser_Parse(t *testing.T) {
	mockClient := &MockClient{}
	type fields struct {
		commands []interfaces.Command
	}
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "upload - failure - does not exist",
			fields: fields{
				commands: []interfaces.Command{
					NewUploadCommand(mockClient),
				},
			},
			args: args{
				args: []string{"upload", "-filename", "doesNotExist.mp3"},
			},
			wantErr: true, // error = open doesNotExist.mp3: no such file or directory
		},
		{
			name: "upload - success - uploaded",
			fields: fields{
				commands: []interfaces.Command{
					NewUploadCommand(mockClient),
				},
			},
			args: args{
				args: []string{"upload", "-filename", "testdata/exists.mp3"},
			},
			wantErr: false,
		},
		{
			name: "get - failure - id does not exist",
			fields: fields{
				commands: []interfaces.Command{
					NewGetCommand(mockClient),
				},
			},
			args: args{
				args: []string{"get", "-id", "567"},
			},
			wantErr: true, // error = audiofile id does not exist
		},
		{
			name: "get - success - requested",
			fields: fields{
				commands: []interfaces.Command{
					NewGetCommand(mockClient),
				},
			},
			args: args{
				args: []string{"get", "-id", "123"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				commands: tt.fields.commands,
			}
			if err := p.Parse(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("Parser.Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
