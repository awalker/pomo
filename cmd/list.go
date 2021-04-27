package cmd

import (
	"fmt"
	"pomo/timers"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	listCmd.Flags().BoolVarP(&listHeaders, "header", "", true, "Prints headers and seperators when listing values")
	listCmd.Flags().BoolVarP(&listTimers, "timers", "t", true, "Prints timer names")
	listCmd.Flags().BoolVarP(&listLabels, "labels", "l", true, "Prints label names")
	listCmd.Flags().BoolVarP(&listCompleted, "history", "", false, "Prints all the completed timers we have tracked")
	listCmd.Flags().BoolVarP(&showActive, "active", "a", true, "Show the active timer if it exists")
	rootCmd.AddCommand(listCmd)
}

var (
	showActive    bool
	listTimers    bool
	listLabels    bool
	listCompleted bool
	listHeaders   bool
	listCmd       = &cobra.Command{
		Use:   "list",
		Short: "List timers and labels",
		Long:  `List all timers and labels.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			timers, err := timers.Load()

			if err != nil {
				return err
			}

			if listHeaders && showActive {
				if a := timers.Active; a != nil {
					t := timers.FindTemplate(a.Template)
					if timers.Paused {
						fmt.Println("Paused: ", t.Name, a.Label)
					} else {
						d := time.Until(*a.Ends)
						fmt.Println("Active: ", t.Name, a.Label, d, "left")
					}
				}
			}

			if listHeaders && listTimers {
				fmt.Println("")
				fmt.Println("Timer Names:")
			}
			if listTimers {
				for _, t := range timers.Templates {
					fmt.Println(t.Name)
				}
			}
			if listHeaders && listLabels {
				fmt.Println("")
				fmt.Println("Timer Labels:")
			}
			if listLabels {
				for _, lbl := range timers.Labels {
					fmt.Println(lbl)
				}
			}
			if listHeaders && listCompleted {
				fmt.Println("")
				fmt.Println("History:")
			}
			if listCompleted {
				for _, c := range timers.Completed {
					t := timers.FindTemplate(c.Template)
					fmt.Println(t.Name, c.Ends)
				}
			}

			return err
		},
	}
)
