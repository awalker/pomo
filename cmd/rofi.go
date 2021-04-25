package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(rofiCmd)
}

var (
	rofiCmd = &cobra.Command{
		Use:   "rofi",
		Short: "List timers and labels",
		Long:  `List all timers and labels.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
		},
	}
)
