package cmd

import (
	"fmt"
	"pomo/timers"

	"github.com/spf13/cobra"
)

func init() {
	// startCmd.Flags().String
	rootCmd.AddCommand(stopCmd)
}

var (
	stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "Cancels the active timer",
		Long:  `Stop and discard the active timer.`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			timers, err := timers.Load()
			if err != nil {
				return err
			}
			if timers.Active != nil {
				timers.Active = nil
				timers.SaveData()
			} else {
				fmt.Println("No Active Timer")
			}

			return err
		},
	}
)
