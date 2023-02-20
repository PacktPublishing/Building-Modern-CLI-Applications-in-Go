package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func process() {
	fmt.Println("Caller group id:", os.Getegid())
	fmt.Println("Caller user id:", os.Geteuid())
	fmt.Println("Process id of caller", os.Getpid())

	cmd := exec.Command(filepath.Join(os.Getenv("GOPATH"), "bin", "sleep"))
	fmt.Println("running sleep for 1 second...")
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	fmt.Println("Process id of sleep", cmd.Process.Pid)
	this, err := os.FindProcess(cmd.Process.Pid)
	if err != nil {
		fmt.Println("unable to find process with id: ", cmd.Process.Pid)
	}
	processState, err := this.Wait()
	if err != nil {
		panic(err)
	}
	if processState.Exited() && processState.Success() {
		fmt.Println("Sleep process ran successfully with exit code: ", processState.ExitCode())
	} else {
		fmt.Println("Sleep process failed with exit code: ", processState.ExitCode())
	}
	fmt.Println(processState.String())
}
