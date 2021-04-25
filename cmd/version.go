package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print version number",
		Long:  `Prints the version numbers.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Pomo version xxx -- HEAD")
		},
	}
)
