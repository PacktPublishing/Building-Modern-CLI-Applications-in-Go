package cmd

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestDelete(t *testing.T) {
	ConfigureTest()
	expectedID := "456"
	b := bytes.NewBufferString("")
	rootCmd.SetOut(b)
	rootCmd.SetArgs([]string{"delete", "--id", expectedID})
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("err: ", err)
	}
	expected := fmt.Sprintf("\U00002705 Successfully deleted audiofile (%s)!\n", expectedID)
	actualBytes, err := io.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	actual := string(actualBytes)
	if !(string(actualBytes) == expected) {
		t.Fatalf("expected \"%s\" got \"%s\"", expected, actual)
	}
}
