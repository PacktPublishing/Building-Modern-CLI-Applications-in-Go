package examples

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func Timeout() {
	errChan := make(chan error, 1)
	cmd := exec.Command(filepath.Join(os.Getenv("GOPATH"), "bin", "timeout"))
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	go func() {
		errChan <- cmd.Wait()
	}()
	select {
	case <-time.After(time.Second * 10):
		fmt.Println("timeout command timed out")
		return
	case err := <-errChan:
		if err != nil {
			fmt.Println("timeout error:", err)
		}
	}
}
