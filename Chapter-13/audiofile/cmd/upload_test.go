//go:build !int

package cmd

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestUpload(t *testing.T) {
	ConfigureTest()
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"upload", "--filename", "list.go"})
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("err: ", err)
	}
	expected := "123"
	actualBytes, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	actual := string(actualBytes)
	if !(string(actualBytes) == expected) {
		t.Fatalf("expected \"%s\" got \"%s\"", expected, actual)
	}
}
