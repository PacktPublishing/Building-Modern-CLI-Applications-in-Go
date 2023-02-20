package cmd

import (
	"encoding/json"

	"github.com/marianina8/audiofile/models"
	"github.com/spf13/cobra"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play audio file by id",
	Long:  `Play audio file by id`,
	RunE: func(cmd *cobra.Command, args []string) error {
		b, err := getAudioByID(cmd)
		if err != nil {
			return err
		}
		audio := models.Audio{}
		err = json.Unmarshal(b, &audio)
		if err != nil {
			return err
		}
		return play(audio.Path)
	},
}

func init() {
	playCmd.Flags().String("id", "", "audiofile id")
	rootCmd.AddCommand(playCmd)
}

// var playCmd = &cobra.Command{
// 	Use:   "play",
// 	Short: "Play audio file by id",
// 	Long:  `Play audio file by id`,
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		b, err := getAudioByID(cmd)
// 		if err != nil {
// 			return err
// 		}
// 		audio := models.Audio{}
// 		err = json.Unmarshal(b, &audio)
// 		if err != nil {
// 			return err
// 		}
// 		switch runtime.GOOS {
// 		case "darwin":
// 			darwinPlay(audio.Path)
// 			return nil
// 		case "windows":
// 			windowsPlay(audio.Path)
// 			return nil
// 		case "linux":
// 			linuxPlay(audio.Path)
// 			return nil
// 		default:
// 			fmt.Println(`Your operating system isn't supported for playing music yet.
// 			Feel free to implement your additional use case!`)
// 		}
// 		return nil
// 	},
// }

// func darwinPlay(audiofilePath string) {
// 	cmd := exec.Command("afplay", audiofilePath)
// 	if err := cmd.Start(); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("enjoy the music!")
// 	err := cmd.Wait()
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func windowsPlay(audiofilePath string) {
// 	cmd := exec.Command("cmd", "/C", "start", audiofilePath)
// 	if err := cmd.Start(); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("enjoy the music!")
// 	err := cmd.Wait()
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func linuxPlay(audiofilePath string) {
// 	cmd := exec.Command("aplay", audiofilePath)
// 	if err := cmd.Start(); err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("enjoy the music!")
// 	err := cmd.Wait()
// 	if err != nil {
// 		panic(err)
// 	}
// }
