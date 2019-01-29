package cmd

import (
	"fmt"
	"os"

	"github.com/RyanTKing/dia/internal/editor"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dia [arguments] [file ...]",
	Short: "dia is a quick and extensible text editor",
	Run: func(cmd *cobra.Command, args []string) {
		state, err := editor.NewState()
		if err != nil {
			die(err)
		}
		defer state.Destroy()

		if err := state.StartLoop(); err != nil {
			die(err)
		}
	},
}

// Execute runs the main command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func die(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}
