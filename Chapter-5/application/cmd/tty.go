/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// ttyCmd represents the tty command
var ttyCmd = &cobra.Command{
	Use: "tty",
	Run: func(cmd *cobra.Command, args []string) {
		if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) != 0 {
			fmt.Println("terminal")
		} else {
			fmt.Println("not a terminal")
		}
	},
}

func init() {
	rootCmd.AddCommand(ttyCmd)
}
