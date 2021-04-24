package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"pomo/rofi"
	"pomo/server"
	"pomo/timers"
)

var rawConfigDir string = "$HOME/.config/pomo"
var serverUrl string = ""

func list() error {
	rofi.Message("Pomo - pomodoro timers")
	fmt.Println("create")
	fmt.Println("pause")
	fmt.Println("start")
	fmt.Println("stop")
	return nil
}

func start() error {
	fmt.Println("do start")
	choice, err := rofi.Dmenu("pomo start> ", "", "standard", "standard with label", "short break", "long break")
	if err != nil {
		return err
	}
	fmt.Println(choice)
	return nil
}

func stop() error {
	return nil
}

func create() error {
	// fmt.Println(rofi("", "prompt", "Label"))
	// rofiMessage("Create a new label")
	// fmt.Println("create")
	label, err := rofi.Input("Enter Label > ")
	if err != nil {
		return err
	}
	fmt.Println(label)
	return nil
}

func rofiMode() error {
	choice, err := rofi.Dmenu("pomo> ", "No timer currently running.", "start", "stop", "pause", "list")
	if err != nil {
		return err
	}
	switch choice {
	case "start":
		err = start()
	case "stop":
		err = stop()
	case "pause":
	case "list":
		err = list()
	}
	return err
}

func main() {
	// Prepare flags
	const rawConfigDirHelp = "The directory to store config data"
	const serverHelp = "url to server"
	flag.StringVar(&rawConfigDir, "config", rawConfigDir, rawConfigDirHelp)
	flag.StringVar(&rawConfigDir, "c", rawConfigDir, rawConfigDirHelp+" <shortcut>")
	flag.StringVar(&serverUrl, "server", serverUrl, serverHelp)

	// Server sub-command flags
	serverFlags := flag.NewFlagSet("server", flag.ExitOnError)
	detach := false
	serverFlags.BoolVar(&detach, "d", detach, "Detach from tty. Daemonize.")

	// Map sub-commands
	flagsMap := make(map[string]*flag.FlagSet)
	flagsMap["server"] = serverFlags

	// Parse all flags
	flag.Parse()
	args := flag.Args()
	cmd := "list"
	if len(args) != 0 {
		cmd = args[0]
		args = args[1:]
		if flags, found := flagsMap[cmd]; found {
			flags.Parse(args)
			args = flags.Args()
		}
	}

	// TODO: if serverUrl is blank, check if server is running on local machine
	// If so, set serverUrl to correct value.

	// Load data if needed
	if serverUrl == "" {
		configDir := os.ExpandEnv(rawConfigDir)
		if err := os.MkdirAll(configDir, 0700); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		timers, err := timers.Load(configDir)
		if err != nil {
			log.Fatal(err)
		}
		_ = timers
		fmt.Println(timers)
	}

	var err error

	// Execute sub-commands
	switch cmd {
	case "list":
		err = list()
	case "start":
		err = start()
	case "stop":
		err = stop()
	case "create":
		err = create()
	case "rofi":
		if len(args) > 0 {
			rofiCmd := args[0]
			fmt.Println(args)
			if rofiCmd == "start" {
				fmt.Println("start standard")
				fmt.Println("start <b>label</b>")
			}
		} else {
			err = list()
		}
	case "dmenu":
		err = rofiMode()
	case "server":
		err = server.Start()
		if detach && err == nil {
			err = server.Detach()
		}
	default:
		fmt.Printf("%s not found\n", cmd)
	}

	// Final error handling and clean up
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
