/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// piperCmd represents the piper command
var piperCmd = &cobra.Command{
	Use: "piper",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("piper called")
		reader := bufio.NewReader(os.Stdin)
		s, _ := reader.ReadString('\n')
		fmt.Printf("piped in: %s\n", s)
	},
}

func init() {
	rootCmd.AddCommand(piperCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// piperCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// piperCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
