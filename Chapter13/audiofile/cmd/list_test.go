//go:build !int

package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/marianina8/audiofile/models"
)

func TestList(t *testing.T) {
	ConfigureTest()
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"list", "--json"})
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("err: ", err)
	}
	actualBytes, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	expectedBytes, err := os.ReadFile("./testfiles/list.json")
	if err != nil {
		t.Fatal(err)
	}
	var audioList1, audioList2 models.AudioList
	json.Unmarshal(actualBytes, &audioList1)
	json.Unmarshal(expectedBytes, &audioList2)
	if len(audioList1) != len(audioList2) {
		t.Fatalf("expected length of list \"%d\" got \"%d\"", len(audioList2), len(audioList1))
	}
}
