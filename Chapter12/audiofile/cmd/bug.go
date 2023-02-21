package cmd

import (
	"bytes"
	"fmt"
	"net/url"

	"github.com/marianina8/audiofile/utils"
	"github.com/spf13/cobra"
)

// bugCmd represents the bug command
var bugCmd = &cobra.Command{
	Use:     "bug",
	Short:   "Submit a bug",
	Long:    "Bug opens the default browser to start a bug report which will include useful system information.",
	Example: `audiofile bug`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return fmt.Errorf("too many arguments")
		}
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("**Audiofile version**\n%s\n\n", utils.Version()))
		buf.WriteString(description)
		buf.WriteString(toReproduce)
		buf.WriteString(expectedBehavior)
		buf.WriteString(additionalDetails)

		body := buf.String()
		url := "https://github.com/marianina8/audiofile/issues/new?title=Bug Report&body=" + url.QueryEscape(body)
		// we print if the browser fails to open
		if !openBrowser(url) {
			fmt.Print("Please file a new issue at https://github.com/marianina8/audiofile/issues/new using this template:\n\n")
			fmt.Print(body)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(bugCmd)
}

const description = `**Description**
A clear description of the bug encountered.

`

const toReproduce = `**To reproduce**
Steps to reproduce the bug.

`
const expectedBehavior = `**Expected behavior**
Expected behavior.

`

const additionalDetails = `**Additional details**
Any other useful data to share.

`
