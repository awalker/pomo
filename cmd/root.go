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
	configFile    string
	serverUrl     string
	dataDirectory string
)

const (
	DATA                           = "data.directory"
	AUTOSTART_BREAKS               = "autostart.breaks"
	AUTOSTART_WORK                 = "autostart.work"
	SHORT_BREAKS_BEFORE_LONG_BREAK = "default.short_breaks_before_long_break"
	GOAL                           = "default.goal"
	DEFAULT_BREAKS_LONG_NAME       = "default.breaks.long.name"
	DEFAULT_BREAKS_SHORT_NAME      = "default.breaks.short.name"
	DEFAULT_WORK_NAME              = "default.work.name"
	CONFIG                         = "config"
	CLIENT_SERVER_URL              = "client.server.url"
	SERVER_PORT                    = "server.port"
)

func init() {
	viper.SetEnvPrefix("pomo")
	cobra.OnInitialize(initConfig)

	defaultDataDirectory := os.ExpandEnv("$HOME/.config/pomo")
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "",
		"Path to the configuration file.")
	rootCmd.PersistentFlags().StringVarP(&serverUrl, "server", "s", "",
		"Url to server.")
	rootCmd.PersistentFlags().StringVarP(&dataDirectory, "data", "",
		defaultDataDirectory, "Path to a directory where pomo data is stored")

	// Could read config from env I guess? Otherwise kinda pointless
	viper.BindEnv(CONFIG)
	viper.BindPFlag(CLIENT_SERVER_URL, rootCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag(DATA, rootCmd.PersistentFlags().Lookup("data"))

	// Basic settings (Not bound to flags)
	viper.SetDefault(AUTOSTART_BREAKS, true)
	viper.SetDefault(AUTOSTART_WORK, false)
	viper.SetDefault(SHORT_BREAKS_BEFORE_LONG_BREAK, 4)
	viper.SetDefault(GOAL, 8)
	viper.SetDefault(DEFAULT_BREAKS_LONG_NAME, "Long Break")
	viper.SetDefault(DEFAULT_BREAKS_SHORT_NAME, "Short Break")
	viper.SetDefault(DEFAULT_WORK_NAME, "Work")

}

func initConfig() {
	if configFile == "" {
		configFile = viper.GetString(CONFIG)
	}
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
		// TODO: This is for debugging. Remove before production
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		fmt.Println("viper.get server.port", viper.GetInt(SERVER_PORT))
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
