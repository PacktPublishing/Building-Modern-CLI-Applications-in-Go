package main

import (
	"github.com/marianina8/audiofile/cmd"
	"github.com/marianina8/audiofile/utils"
)

func main() {
	cmd.Configure()
	utils.InitCLILogger()
	cmd.Execute()
}
