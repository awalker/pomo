package cmd

import (
	"errors"
	"pomo/timers"

	"github.com/spf13/cobra"
)

func init() {
	// startCmd.Flags().String
	rootCmd.AddCommand(startCmd)
}

var (
	startCmd = &cobra.Command{
		Use:   "start [timer] [label]",
		Short: "Start a timer",
		Long:  `Start a timer or restart a pauser timer`,
		Args: func(cmd *cobra.Command, args []string) error {
			// TODO: Check for a paused timer.
			if len(args) < 1 {
				return errors.New("Requires a timer name")
			}
			// TODO: Validate the timer name

			// TODO: Check labels
			/*if myapp.IsValidColor(args[0]) {
				return nil
			}*/
			// return fmt.Errorf("invalid color specified: %s", args[0])
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			timers, err := timers.Load()
			timers.Paused = false
			if err != nil {
				return err
			}
			label := ""
			if len(args) > 1 {
				label = args[1]
			}

			err = timers.Start(args[0], label)
			if err == nil {
				err = timers.SaveData()
			}

			return err
		},
	}
)
