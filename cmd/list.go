package cmd

import (
	"pomo/timers"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "List timers and labels",
		Long:  `List all timers and labels.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			timers, err := timers.Load()

			_ = timers

			return err
		},
	}
)
