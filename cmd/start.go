package cmd

import (
	"pomo/timers"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var (
	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start a timer",
		Long:  `Start a timer`,
		RunE: func(cmd *cobra.Command, args []string) error {
			timers, err := timers.Load()
			if err != nil {
				return err
			}

			err = timers.Start("Work")
			if err == nil {
				timers.SaveData()
			}

			return err
		},
	}
)
