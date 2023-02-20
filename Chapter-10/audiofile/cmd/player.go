package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"image"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/marianina8/audiofile/models"
	"github.com/mitchellh/go-ps"
	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/align"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/tcell"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/button"
	"github.com/mum4k/termdash/widgets/segmentdisplay"
	"github.com/mum4k/termdash/widgets/text"
	"github.com/mum4k/termdash/widgets/textinput"

	"github.com/spf13/cobra"
)

var audiofileID = -1
var pID = 0

func newStopButton() (*button.Button, error) {
	stopButton, err := button.New("Stop", func() error {
		go func() error {
			stopTheMusic()
			pID = 0
			return nil
		}()
		return nil
	},
		button.FillColor(cell.ColorNumber(9)),
		button.GlobalKey('s'),
	)
	if err != nil {
		return stopButton, fmt.Errorf("%v", err)
	}
	return stopButton, nil
}

func newPlayButton(audioList *models.AudioList, playID <-chan int) (*button.Button, error) {
	playButton, err := button.New("Play", func() error {
		stopTheMusic()
		go func() {
			if audiofileID <= len(*audioList)-1 && audiofileID >= 0 {
				pID, _ = play((*audioList)[audiofileID].Path, false, true)
			}
		}()
		return nil
	},
		button.FillColor(cell.ColorNumber(220)),
		button.GlobalKey('p'),
	)
	if err != nil {
		return playButton, fmt.Errorf("%v", err)
	}
	return playButton, nil
}

func newTextInput(audioList *models.AudioList, updatedID chan<- int, updateText, errorText chan<- string) *textinput.TextInput {
	input, _ := textinput.New(
		textinput.Label("Enter id of song: ", cell.FgColor(cell.ColorNumber(33))),
		textinput.MaxWidthCells(20),
		textinput.OnSubmit(func(text string) error {
			number, _ := strconv.Atoi(text)
			audioListLength := len(*audioList)
			if number >= audioListLength || number < 0 {
				errorText <- fmt.Errorf("Id doesn't exist, Id needs to be between 0 and %d", audioListLength-1).Error()
			} else {
				updatedID <- number
				if (*audioList)[number].Metadata.Tags.Title != "" && (*audioList)[number].Metadata.Tags.Artist != "" {
					updateText <- "Enjoy the music..." + (*audioList)[number].Metadata.Tags.Title + " by " + (*audioList)[number].Metadata.Tags.Artist
				} else {
					updateText <- "Enjoy the music..." + filepath.Base((*audioList)[number].Path)
				}
				audiofileID = number
				errorText <- " "
			}
			return nil
		}),
		textinput.ClearOnSubmit(),
	)
	return input
}

func newErrorText(errors chan string) *text.Text {
	errorText, _ := text.New(text.WrapAtWords(), text.RollContent())
	go func() {
		for {
			errText := <-errors
			errorText.Write(errText, text.WriteReplace(), text.WriteCellOpts(cell.FgColor(cell.ColorRed)))
		}
	}()
	return errorText
}

func newMetadataDisplay(t terminalapi.Terminal, audioList *models.AudioList, updatedID <-chan int) (*text.Text, error) {
	metadata, err := text.New(text.WrapAtWords(), text.RollContent())
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			id := <-updatedID
			if id >= len(*audioList) || id < 0 {
				continue
			}
			idString := strconv.Itoa(id)
			metadata.Write("[id="+idString+"]\n\n", text.WriteReplace())
			textExists := false
			if (*audioList)[id].Metadata.Tags.Title != "" {
				textExists = true
				metadata.Write("Title: ", text.WriteCellOpts(cell.FgColor(cell.ColorNumber(33))))
				metadata.Write((*audioList)[id].Metadata.Tags.Title)
				metadata.Write("\n")
			}
			if (*audioList)[id].Metadata.Tags.AlbumArtist != "" {
				textExists = true
				metadata.Write("AlbumArtist: ", text.WriteCellOpts(cell.FgColor(cell.ColorNumber(33))))
				metadata.Write((*audioList)[id].Metadata.Tags.AlbumArtist)
				metadata.Write("\n")
			}
			if (*audioList)[id].Metadata.Tags.Artist != "" {
				textExists = true
				metadata.Write("Artist: ", text.WriteCellOpts(cell.FgColor(cell.ColorNumber(33))))
				metadata.Write((*audioList)[id].Metadata.Tags.Artist)
				metadata.Write("\n")
			}
			if (*audioList)[id].Metadata.Tags.Album != "" {
				textExists = true
				metadata.Write("Album: ", text.WriteCellOpts(cell.FgColor(cell.ColorNumber(33))))
				metadata.Write((*audioList)[id].Metadata.Tags.Album)
				metadata.Write("\n")
			}
			if (*audioList)[id].Metadata.Tags.Composer != "" {
				textExists = true
				metadata.Write("Composer: ", text.WriteCellOpts(cell.FgColor(cell.ColorNumber(33))))
				metadata.Write((*audioList)[id].Metadata.Tags.Composer)
				metadata.Write("\n")
			}
			if (*audioList)[id].Metadata.Tags.Genre != "" {
				textExists = true
				metadata.Write("Genre: ", text.WriteCellOpts(cell.FgColor(cell.ColorNumber(33))))
				metadata.Write((*audioList)[id].Metadata.Tags.Genre)
				metadata.Write("\n")
			}
			if (*audioList)[id].Metadata.Tags.Year > 0 {
				textExists = true
				metadata.Write("Year: ", text.WriteCellOpts(cell.FgColor(cell.ColorNumber(33))))
				metadata.Write(strconv.Itoa((*audioList)[id].Metadata.Tags.Year))
				metadata.Write("\n")
			}
			if (*audioList)[id].Metadata.Transcript != "" {
				textExists = true
				metadata.Write("Transcript: ", text.WriteCellOpts(cell.FgColor(cell.ColorNumber(33))))
				metadata.Write((*audioList)[id].Metadata.Transcript)
				metadata.Write("\n")
			}
			if (*audioList)[id].Metadata.Tags.Lyrics != "" {
				textExists = true
				metadata.Write("Lyrics: ", text.WriteCellOpts(cell.FgColor(cell.ColorNumber(33))))
				metadata.Write((*audioList)[id].Metadata.Tags.Lyrics)
				metadata.Write("\n")
			}
			if strings.TrimSpace(strings.TrimSuffix((*audioList)[id].Metadata.Tags.Comment, "\n")) != "" {
				textExists = true
				metadata.Write("Comment: ", text.WriteCellOpts(cell.FgColor(cell.ColorNumber(33))))
				metadata.Write((*audioList)[id].Metadata.Tags.Comment)
				metadata.Write("\n")
			}
			if !textExists {
				metadata.Write("No metadata available.")
			}
		}
	}()
	return metadata, err
}

func newLibraryContent(audioList *models.AudioList) (*text.Text, error) {
	libraryContent, err := text.New(text.RollContent(), text.WrapAtWords())
	if err != nil {
		panic(err)
	}
	for i, audiofile := range *audioList {
		if audiofile.Metadata.Tags.Title != "" && audiofile.Metadata.Tags.Artist != "" {
			libraryContent.Write(fmt.Sprintf("[id=%d] %s by %s\n", i, audiofile.Metadata.Tags.Title, audiofile.Metadata.Tags.Artist))
		} else {
			libraryContent.Write(fmt.Sprintf("[id=%d] %s\n", i, filepath.Base(audiofile.Path)))
		}
	}
	return libraryContent, nil
}

// playerCmd represents the player command
var playerCmd = &cobra.Command{
	Use:   "player",
	Short: "Launch player dashboard",
	Long: `Launches a terminal player dashboard where you can view 
	all uploaded songs metadata, select and play from a dashboard.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		audioList := &models.AudioList{}
		b, err := callList(false)
		if err != nil {
			return err
		}
		json.Unmarshal(b, &audioList)

		t, err := tcell.New(tcell.ColorMode(terminalapi.ColorMode256))
		if err != nil {
			return fmt.Errorf("error creating new tcell: %s", err.Error())
		}
		defer t.Close()

		ctx, cancel := context.WithCancel(context.Background())

		updatedID := make(chan int)
		updateText := make(chan string)
		errors := make(chan string)
		playID := make(chan int)
		input := newTextInput(audioList, updatedID, updateText, errors)

		errorText := newErrorText(errors)

		playButton, err := newPlayButton(audioList, playID)
		if err != nil {
			cancel()
			return err
		}

		stopButton, err := newStopButton()
		if err != nil {
			cancel()
			return err
		}

		libraryContent, err := newLibraryContent(audioList)
		if err != nil {
			cancel()
			return err
		}

		metadataDisplay, err := newMetadataDisplay(t, audioList, updatedID)
		if err != nil {
			cancel()
			return err
		}

		streamingTitle, err := newStreamingTitle(ctx, t, updateText)
		if err != nil {
			cancel()
			return err
		}

		c, err := container.New(
			t,
			container.SplitVertical(
				container.Left(
					container.Border(linestyle.Light),
					container.BorderTitle("Music Library"),
					container.PlaceWidget(libraryContent),
				),
				container.Right(
					container.SplitHorizontal(
						container.Top(
							container.SplitHorizontal(
								container.Top(
									container.SplitHorizontal(
										container.Top(
											container.PlaceWidget(input),
										),
										container.Bottom(
											container.PlaceWidget(errorText),
											container.AlignHorizontal(align.HorizontalCenter),
											container.AlignVertical(align.VerticalMiddle),
										),
										container.SplitPercent(60),
									),
								),
								container.Bottom(
									container.BorderTitle("Selected song"),
									container.Border(linestyle.Light),
									container.PlaceWidget(streamingTitle),
								),
								container.SplitPercent(30),
							),
						),
						container.Bottom(
							container.SplitHorizontal(
								container.Top(
									container.Border(linestyle.Light),
									container.BorderTitle("Metadata"),
									container.PlaceWidget(metadataDisplay),
								),
								container.Bottom(
									container.SplitVertical(
										container.Left(
											container.PlaceWidget(playButton),
										),
										container.Right(
											container.PlaceWidget(stopButton),
										),
									)),
								container.SplitPercent(80),
							),
						),
						container.SplitPercent(30),
					),
				),
			),
			container.Border(linestyle.Light),
			container.BorderTitle("Press Q to quit"),
		)
		if err != nil {
			cancel()
			return err
		}
		quitter := func(k *terminalapi.Keyboard) {
			if k.Key == 'q' || k.Key == 'Q' {
				stopTheMusic()
				cancel()
			}
		}

		if err := termdash.Run(ctx, t, c, termdash.KeyboardSubscriber(quitter), termdash.RedrawInterval(100*time.Millisecond)); err != nil {
			panic(err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(playerCmd)
}

// newStreamingTitle creates a new SegmentDisplay that initially shows the
// Termdash name. Shows any text that is sent over the channel.
func newStreamingTitle(ctx context.Context, t terminalapi.Terminal, updateText <-chan string) (*segmentdisplay.SegmentDisplay, error) {
	sd, err := segmentdisplay.New()
	if err != nil {
		return nil, err
	}

	colors := []cell.Color{
		cell.ColorNumber(33),
		cell.ColorRed,
		cell.ColorYellow,
		cell.ColorNumber(33),
		cell.ColorGreen,
		cell.ColorRed,
		cell.ColorGreen,
		cell.ColorRed,
	}

	text := "Audiofile"
	step := 0

	go func() {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		capacity := 0
		termSize := t.Size()
		for {
			select {
			case <-ticker.C:
				if capacity == 0 {
					// The segment display only knows its capacity after both
					// text size and terminal size are known.
					capacity = sd.Capacity()
				}
				if t.Size().Eq(image.ZP) || !t.Size().Eq(termSize) {
					// Update the capacity initially the first time the
					// terminal reports a non-zero size and then every time the
					// terminal resizes.
					//
					// This is better than updating the capacity on every
					// iteration since that leads to edge cases - segment
					// display capacity depends on the length of text and here
					// we are trying to adjust the text length to the capacity.
					termSize = t.Size()
					capacity = sd.Capacity()
				}

				state := textState(text, capacity, step)
				var chunks []*segmentdisplay.TextChunk
				for i := 0; i < capacity; i++ {
					if i >= len(state) {
						break
					}

					color := colors[i%len(colors)]
					chunks = append(chunks, segmentdisplay.NewChunk(
						string(state[i]),
						segmentdisplay.WriteCellOpts(cell.FgColor(color)),
					))
				}
				if len(chunks) == 0 {
					continue
				}
				if err := sd.Write(chunks); err != nil {
					panic(err)
				}
				step++

			case t := <-updateText:
				text = t
				sd.Reset()
				step = 0

			case <-ctx.Done():
				return
			}
		}
	}()
	return sd, nil
}

func rotate(inputs []rune, step int) []rune {
	return append(inputs[step:], inputs[:step]...)
}

// newRollText creates a new Text widget that displays rolling text.
func newRollText(ctx context.Context) (*text.Text, error) {
	t, err := text.New(text.RollContent())
	if err != nil {
		return nil, err
	}

	i := 0
	go periodic(ctx, 1*time.Second, func() error {
		if err := t.Write(fmt.Sprintf("Writing line %d.\n", i), text.WriteCellOpts(cell.FgColor(cell.ColorNumber(142)))); err != nil {
			return err
		}
		i++
		return nil
	})
	return t, nil
}

// periodic executes the provided closure periodically every interval.
// Exits when the context expires.
func periodic(ctx context.Context, interval time.Duration, fn func() error) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if err := fn(); err != nil {
				panic(err)
			}
		case <-ctx.Done():
			return
		}
	}
}

// textState creates a rotated state for the text we are displaying.
func textState(text string, capacity, step int) []rune {
	if capacity == 0 {
		return nil
	}

	var state []rune
	for i := 0; i < capacity; i++ {
		state = append(state, ' ')
	}
	state = append(state, []rune(text)...)
	step = step % len(state)
	return rotateRunes(state, step)
}

// rotateRunes returns a new slice with inputs rotated by step.
// I.e. for a step of one:
//
//	inputs[0] -> inputs[len(inputs)-1]
//	inputs[1] -> inputs[0]
//
// And so on.
func rotateRunes(inputs []rune, step int) []rune {
	return append(inputs[step:], inputs[:step]...)
}

func stopTheMusic() {
	processes, _ := ps.Processes()
	var playerExecutable string
	switch runtime.GOOS {
	case "windows":
		playerExecutable = "start"
	case "linux":
		playerExecutable = "aplay"
	case "darwin":
		playerExecutable = "afplay"
	default:
		playerExecutable = "unknown"
	}
	for _, p := range processes {
		if p.Executable() == playerExecutable {
			proc, _ := os.FindProcess(p.Pid())
			if proc != nil {
				proc.Kill()
			}
		}
	}
}
