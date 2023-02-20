package cmd

import (
	"github.com/spf13/cobra"
)

var (
	Filename string
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload [audio|video] [-f|--filename] <filename>",
	Short: "upload an audio or video file",
	Long: `This command allows you to upload either an audio or video file for metadata extraction.
	To pass in a filename, use the -f or --filename flag followed by the path of the file.
	
	Examples:
	./audiofile-cli upload audio -f audio/beatdoctor.mp3
	./audiofile-cli upload video --filename video/musicvideo.mp4`,
	SuggestFor: []string{"add"},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		return []string{"audio", "video"}, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	uploadCmd.PersistentFlags().StringVarP(&Filename, "filename", "f", "", "file to upload")
	uploadCmd.MarkPersistentFlagRequired("filename")

	rootCmd.AddCommand(uploadCmd)
}

/************************ DEFAULT GENERATED CODE BY COBRA-CLI **************************

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("upload called")
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// uploadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
****************************************************************************************/
