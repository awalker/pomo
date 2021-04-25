package cmd

import (
	"log"
	"pomo/rofi"

	"github.com/spf13/cobra"
)

func rofiMode() error {
	choice, err := rofi.Dmenu("pomo> ", "No timer currently running.", "start", "stop", "pause", "list")
	if err != nil {
		return err
	}
	switch choice {
	case "start":
		// err = start()
	case "stop":
		// err = stop()
	case "pause":
	case "list":
		// err = list()
	}
	return err
}

func init() {
	rootCmd.AddCommand(dmenuCmd)
}

var (
	dmenuCmd = &cobra.Command{
		Use:   "dmenu",
		Short: "Run rofi in dmenu mode",
		Long:  `Run rofi in dmenu mode`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := rofiMode(); err != nil {
				log.Fatal(err)
			}
		},
	}
)
