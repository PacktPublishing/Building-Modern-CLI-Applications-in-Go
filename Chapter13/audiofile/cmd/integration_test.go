//go:build int && pro

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/marianina8/audiofile/models"
	"github.com/marianina8/audiofile/utils"
	"github.com/spf13/viper"

	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"
)

func ConfigureTest() {
	getClient = &http.Client{
		Timeout: 15 * time.Second,
	}
	viper.SetDefault("cli.hostname", "localhost")
	viper.SetDefault("cli.port", 8000)
	utils.InitCLILogger()
}

func TestWorkflow(t *testing.T) {
	ConfigureTest()
	fmt.Println("*** Testing upload ***")
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"upload", "--filename", "../audio/algorithms.mp3"})
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("err: ", err)
	}
	uploadResponse, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	id := string(uploadResponse)
	fmt.Println("id: ", id)
	if id == "" {
		t.Fatalf("expected id returned")
	}
	fmt.Println("*** Testing get ***")
	rootCmd.SetArgs([]string{"get", "--id", id, "--json"})
	err = rootCmd.Execute()
	if err != nil {
		fmt.Println("err: ", err)
	}
	getResponse, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	var audio models.Audio
	json.Unmarshal(getResponse, &audio)
	if audio.Id != id {
		t.Fatalf("expected matching audiofile returned")
	}

	fmt.Println("*** Testing list ***")
	rootCmd.SetArgs([]string{"list", "--json"})
	err = rootCmd.Execute()
	if err != nil {
		fmt.Println("err: ", err)
	}
	listResponse, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	var audioList models.AudioList
	json.Unmarshal(listResponse, &audioList)
	if len(audioList) < 0 {
		t.Fatalf("expected length of list to be greater than 1, got \"%d\"", len(audioList))
	}
	found := false
	for _, audio := range audioList {
		if audio.Id == id {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected new audio file to be found")
	}
	fmt.Println("*** Testing search ***")
	rootCmd.SetArgs([]string{"search", "--value", "algo"})
	err = rootCmd.Execute()
	if err != nil {
		fmt.Println("err: ", err)
	}
	searchResponse, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	json.Unmarshal(searchResponse, &audioList)
	if len(audioList) < 0 {
		t.Fatalf("expected length of list to be greater than or equal to 1, got \"%d\"", len(audioList))
	}
	found = false
	for _, audio := range audioList {
		if audio.Id == id {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("expected matching audiofile returned from list")
	}
	fmt.Println("*** Testing delete ***")
	rootCmd.SetArgs([]string{"delete", "--id", id})
	err = rootCmd.Execute()
	if err != nil {
		fmt.Println("err: ", err)
	}
	deleteResponse, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(deleteResponse), "Successfully deleted audiofile") || !strings.Contains(string(deleteResponse), id) {
		t.Fatalf("expected audiofile to be successfully deleted")
	}
}
