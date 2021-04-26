package cmd

import (
	"fmt"
	"log"
	"pomo/server"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	serverCmd.Flags().BoolVarP(&detach, "detach", "d", false, "Detach from TTY. Daemonize.")
	serverCmd.Flags().IntVarP(&serverPort, "port", "p", 3966, "Port for listening")
	rootCmd.AddCommand(serverCmd)

	viper.BindPFlag("server.port", serverCmd.Flags().Lookup("port"))
	viper.BindPFlag("server.detach", serverCmd.Flags().Lookup("detach"))
}

var (
	detach     = false
	serverPort int
	serverCmd  = &cobra.Command{
		Use:   "server",
		Short: "Start a pomo server",
		Long:  `Start a pomo server.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("port", viper.GetInt(SERVER_PORT))
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
