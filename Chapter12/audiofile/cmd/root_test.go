package cmd

import (
	"github.com/marianina8/audiofile/utils"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Logger *zap.Logger
var Verbose *zap.Logger

func ConfigureTest() {
	getClient = &ClientMock{}
	viper.SetDefault("cli.hostname", "testHostname")
	viper.SetDefault("cli.port", 8000)
	utils.InitCLILogger()
}
