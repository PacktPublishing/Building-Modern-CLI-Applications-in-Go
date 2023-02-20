package cmd

import (
	"encoding/json"

	"github.com/marianina8/audiofile/models"
	"github.com/marianina8/audiofile/utils"
	"github.com/spf13/cobra"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:     "play",
	Short:   "Play audio file by id",
	Long:    `Play audio file by id using the default audio player for your current system`,
	Example: `./bin/audiofile play --id 45705eba-9342-4952-8cd4-baa2acc25188`,
	RunE: func(cmd *cobra.Command, args []string) error {
		verbose, _ := cmd.Flags().GetBool("verbose")
		b, err := getAudioByID(cmd, verbose)
		if err != nil {
			return err
		}
		audio := models.Audio{}
		err = json.Unmarshal(b, &audio)
		if err != nil {
			return utils.Error("\n  unmarshalling audio struct: %v", err, verbose)
		}
		_, err = play(audio.Path, verbose)
		return err
	},
}

func init() {
	playCmd.Flags().String("id", "", "audiofile id")
	rootCmd.AddCommand(playCmd)
}
