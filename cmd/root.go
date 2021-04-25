package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "pomo",
		Short: "Pomo is a pomodoro timer with server, CLI, and GUI modes",
		Long: `Pomo is a pomodoro timer with server, CLI, and GUI modes
                love by Adam Walker and friends in Go.
                Complete documentation is available at http://...`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
	configFile string
	serverUrl  string
)

func init() {
	defaultConfigFile := os.ExpandEnv("$HOME/.config/pomo/pomo.json")
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", defaultConfigFile, "Path to the configuration file.")
	rootCmd.PersistentFlags().StringVarP(&serverUrl, "server", "s", "", "Url to server.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
