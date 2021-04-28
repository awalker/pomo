package cmd

import (
	"fmt"
	"pomo/timers"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	// startCmd.Flags().String
	rootCmd.AddCommand(pauseCmd)
}

var (
	pauseCmd = &cobra.Command{
		Use:   "pause",
		Short: "Toggles the pause state of the timer",
		Long:  `Toggles the pause state of the timer`,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			timers, err := timers.Load()
			if err != nil {
				return err
			}
			timers.Paused = !timers.Paused
			if timers.Active != nil {
				now := time.Now()
				if timers.Paused {
					timers.Active.Elapsed += now.Sub(*timers.Active.Started)
					timers.Active.Ends = nil
				} else {
					dur := timers.Active.Duration - timers.Active.Elapsed
					if dur > 0 {
						ends := now.Add(dur)
						timers.Active.Ends = &ends
					} else {
						return fmt.Errorf("Negative duration")
					}
				}
			} else {
				fmt.Println("No active timer to pause")
			}
			err = timers.SaveData()

			return err
		},
	}
)
