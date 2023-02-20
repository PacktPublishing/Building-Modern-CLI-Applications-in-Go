/*
Copyright © 2023 Marian Montagnino <mmontagnino@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const (
	storageFile = "storage/result"
)

var (
	value = 0.0
)

func init() {
	value = 0.0
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "calculator",
	Short: "A basic calculator CLI",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
