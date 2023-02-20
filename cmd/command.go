/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// commandCmd represents the command command
var commandCmd = &cobra.Command{
	Use: "command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("command called")
		localFlag, _ := cmd.Flags().GetString("localFlag")
		if localFlag != "" {
			fmt.Printf("localFlag is set to %s\n", localFlag)
		}
	},
}

func init() {
	commandCmd.Flags().String("localFlag", "", "a local string flag")
	commandCmd.PersistentFlags().Bool("persistentFlag", false, "a persistent boolean flag")
	rootCmd.AddCommand(commandCmd)
}
