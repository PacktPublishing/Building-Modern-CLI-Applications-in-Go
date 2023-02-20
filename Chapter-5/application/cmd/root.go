/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "application",
}

func Execute() {
	SetupInterruptHandler()
	SetupStopHandler()
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func SetupInterruptHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	go func() {
		<-c
		fmt.Println("\r- Wake up! Sleep has been interrupted.")
		os.Exit(0)
	}()
}

func SetupStopHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTSTP)
	go func() {
		<-c
		fmt.Println("\r- Wake up! Stopped sleeping.")
		os.Exit(0)
	}()
}
