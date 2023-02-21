package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RootCMD() *cobra.Command {
	return rootCmd
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "audiofile",
	Short: "A command line interface to interact with the Audiofile service",
	Long: `A command line interface allows you to interact with the Audiofile service.
Basic commands include: get, list, and upload.`,
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
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose")
}

func Configure() {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("cli")
	viper.SetConfigType("json")
	viper.ReadInConfig()
	viper.SetDefault("cli.hostname", "localhost")
	viper.SetDefault("cli.port", 8000)
}
