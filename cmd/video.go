package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// videoCmd represents the video command
var videoCmd = &cobra.Command{
	Use:   "video",
	Short: "sets video as the upload type",
	RunE: func(cmd *cobra.Command, args []string) error {
		filename, err := cmd.Flags().GetString("filename")
		if err != nil {
			fmt.Printf("error retrieving filename: %s\n", err.Error())
			return err
		}
		if filename == "" {
			return errors.New("missing filename")
		}
		fmt.Println("uploading video file, ", filename)
		return nil
	},
}

func init() {
	uploadCmd.AddCommand(videoCmd)
}

/************************ DEFAULT GENERATED CODE BY COBRA-CLI **************************
// videoCmd represents the video command
var videoCmd = &cobra.Command{
	Use:   "video",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("video called")
	},
}

func init() {
	rootCmd.AddCommand(videoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// videoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// videoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

****************************************************************************************/
