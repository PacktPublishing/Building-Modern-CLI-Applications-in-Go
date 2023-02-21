//go:build !int

package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestBug(t *testing.T) {
	ConfigureTest()
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"bug", "unexpected"})
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("err: ", err)
	}
	actualBytes, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	expectedBytes, err := os.ReadFile("./testfiles/bug.txt")
	if err != nil {
		t.Fatal(err)
	}
	if strings.TrimSpace(string(actualBytes)) != strings.TrimSpace(string(expectedBytes)) {
		t.Fatal(string(actualBytes), "!=", string(expectedBytes))
	}
}
