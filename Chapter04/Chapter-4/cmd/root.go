package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "audiofile-cli",
	Short: "Audiofile is a metadata extraction tool",
	Long:  `Audiofile extracts metadata from uploaded audio and saves it in storage for on-demand retrieval.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
