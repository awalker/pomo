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
		Short: "todo",
		Long:  `todo`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("todo")
		},
	}
)
