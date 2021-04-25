package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(writeConfigCmd)
}

var (
	writeConfigCmd = &cobra.Command{
		Use:   "write-config",
		Short: "write the config",
		Long:  `Write the config file.`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := viper.WriteConfig(); err != nil {
				panic(fmt.Errorf("Fatal error writing config file: %w \n", err))
			}
		},
	}
)
