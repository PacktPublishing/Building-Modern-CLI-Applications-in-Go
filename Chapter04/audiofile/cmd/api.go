/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"flag"
	"fmt"

	metadataService "github.com/marianina8/audiofile/services/metadata"
	"github.com/spf13/cobra"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start or stop the API required by the CLI",
	Long: `Start or stop the API with the following usage:
	./audiofile api`,
	Run: func(cmd *cobra.Command, args []string) {
		var port int
		flag.IntVar(&port, "p", 8000, "Port for metadata service")
		flag.Parse()
		fmt.Printf("Starting API at http://localhost:%d\nPress Ctrl-C to stop.\n", port)
		metadataService.Run(port)
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
