package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger
var Verbose *zap.Logger

func ensureDir(fileName string) {
	dirName := filepath.Dir(fileName)
	if _, serr := os.Stat(dirName); serr != nil {
		merr := os.MkdirAll(dirName, os.FileMode(0777))
		if merr != nil {
			panic(merr)
		}
	}
}

func createFilesIfNotExists(filenames []string) error {
	for _, filename := range filenames {
		if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
			ensureDir(filename)
			file, err := os.Create(filename)
			if err != nil {
				return err
			}
			defer file.Close()
		}
	}
	return nil
}

func encoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "name",
		CallerKey:      "file",
		StacktraceKey:  "stacktrace",
		EncodeName:     zapcore.FullNameEncoder,
		EncodeTime:     timeEncoder,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func InitCLILogger() {
	var err error
	var cfg zap.Config
	config := viper.GetStringMap("cli.logging")
	configBytes, _ := json.Marshal(config)
	if err := json.Unmarshal(configBytes, &cfg); err != nil {
		panic(err)
	}
	cfg.EncoderConfig = encoderConfig()
	err = createFilesIfNotExists(cfg.OutputPaths)
	if err != nil {
		panic(err)
	}
	cfg.Encoding = "json"
	cfg.Level = zap.NewAtomicLevel()
	Logger, err = cfg.Build()
	if err != nil {
		panic(err)
	}
	cfg.OutputPaths = append(cfg.OutputPaths, "stdout")
	Verbose, err = cfg.Build()
	if err != nil {
		panic(err)
	}
	defer Logger.Sync()
}

func InitAPILogger() {
	var err error
	var cfg zap.Config
	config := viper.GetStringMap("api.logging")
	configBytes, _ := json.Marshal(config)
	if err := json.Unmarshal(configBytes, &cfg); err != nil {
		panic(err)
	}
	cfg.EncoderConfig = encoderConfig()
	err = createFilesIfNotExists(cfg.OutputPaths)
	if err != nil {
		panic(err)
	}
	cfg.OutputPaths = append(cfg.OutputPaths, "stdout")
	Logger, err = cfg.Build()
	if err != nil {
		panic(err)
	}
	defer Logger.Sync()
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func LogRequest(verbose bool, method, path, payload string) {
	if verbose {
		Verbose.Info(fmt.Sprintf("sending request: %s %s %s...\n", method, path, payload))
	} else {
		Logger.Info(fmt.Sprintf("sending request: %s %s %s...\n", path, path, payload))
	}
}

func LogHTTPResponse(verbose bool, resp *http.Response, body []byte) {
	if verbose && resp != nil {
		Verbose.Info(fmt.Sprintf("response status: %s, body: %s", resp.Status, string(body)))
	} else if resp != nil {
		Logger.Info(fmt.Sprintf("response status: %s, body: %s", resp.Status, string(body)))
	}
}
