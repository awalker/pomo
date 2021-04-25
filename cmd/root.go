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

func init() {
	viper.SetEnvPrefix("pomo")
	cobra.OnInitialize(initConfig)

	defaultDataDirectory := os.ExpandEnv("$HOME/.config/pomo")
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "Path to the configuration file.")
	rootCmd.PersistentFlags().StringVarP(&serverUrl, "server", "s", "", "Url to server.")
	rootCmd.PersistentFlags().StringVarP(&dataDirectory, "data", "", defaultDataDirectory, "Path to a directory where pomo data is stored")

	// Could read config from env I guess? Otherwise kinda pointless
	viper.BindEnv("config")
	//viper.BindPFlag("config", rootCmd.PersistentFlags().Lookup("config"))
	viper.BindPFlag("client.server.url", rootCmd.PersistentFlags().Lookup("server"))
	viper.BindPFlag("data", rootCmd.PersistentFlags().Lookup("data"))

}

func initConfig() {
	if configFile == "" {
		configFile = viper.GetString("config")
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
		fmt.Println("viper.get server.port", viper.GetInt("server.port"))
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
