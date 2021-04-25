package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
			cmd.Usage()
		},
	}
	configFile string
	serverUrl  string
)

func init() {
	cobra.OnInitialize(initConfig)

	// defaultConfigFile := os.ExpandEnv("$HOME/.config/pomo/pomo.json")
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "Path to the configuration file.")
	rootCmd.PersistentFlags().StringVarP(&serverUrl, "server", "s", "", "Url to server.")

	// Could read config from env I guess? Otherwise kinda pointless
	viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("server", rootCmd.PersistentFlags().Lookup("server"))
}

func initConfig() {
	if configFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(configFile)
	} else {
		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath("/etc/pomo")
		viper.AddConfigPath(os.ExpandEnv("$HOME/.config/pomo"))
		viper.SetConfigName("pomo")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
