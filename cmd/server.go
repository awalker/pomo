package cmd

import (
	"log"
	"pomo/server"

	"github.com/spf13/cobra"
)

func init() {
	serverCmd.Flags().BoolVarP(&detach, "detach", "d", false, "Detach from TTY. Daemonize.")
	rootCmd.AddCommand(serverCmd)
}

var (
	detach    = false
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Start a pomo server",
		Long:  `Start a pomo server.`,
		Run: func(cmd *cobra.Command, args []string) {
			err := server.Start()
			if err == nil && detach {
				err = server.Detach()
			}
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)
