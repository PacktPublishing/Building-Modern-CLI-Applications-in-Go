/*
Copyright © 2023 Marian Montagnino <mmontagnino@gmail.com>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/marianina8/Chapter-4-Demo/storage"
	"github.com/spf13/cobra"
)

// subtractCmd represents the subtract command
var subtractCmd = &cobra.Command{
	Use:   "subtract number",
	Short: "Subtract value",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 {
			fmt.Println("only accepts a single argument")
			return
		}
		if len(args) == 0 {
			fmt.Println("command requires input value")
			return
		}
		floatVal, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			fmt.Printf("unable to parse input[%s]: %v", args[0], err)
			return
		}
		value = storage.GetValue()
		value -= floatVal
		storage.SetValue(value)
		fmt.Printf("%f\n", value)
	},
}

func init() {
	rootCmd.AddCommand(subtractCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subtractCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subtractCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
