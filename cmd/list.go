package cmd

import (
	"fmt"
	"pomo/timers"
	"time"

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

			if err != nil {
				return err
			}

			if a := timers.Active; a != nil {
				t := timers.FindTemplate(a.Template)
				if timers.Paused {
					fmt.Println("Paused: ", t.Name)
				} else {
					d := time.Until(a.Ends)
					fmt.Println("Active: ", t.Name, d, "left")
				}
			}

			fmt.Println("Timer Names:")
			for _, t := range timers.Templates {
				fmt.Println("\t", t.Name)
			}
			fmt.Println("Timer Labels:")
			for _, lbl := range timers.Labels {
				fmt.Println("\t", lbl)
			}

			return err
		},
	}
)
