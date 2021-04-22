package main

import (
	"flag"
	"fmt"
	"os"
	"pomo/rofi"
)

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

func daemon() error {
	return nil
}

func main() {
	flag.Parse()
	args := flag.Args()
	cmd := "list"
	if len(args) != 0 {
		cmd = args[0]
		args = args[1:]
	}

	configDir := os.ExpandEnv("$HOME/.config/pomo")
	// configFilename := configDir + "/pomo.json"
	if err := os.MkdirAll(configDir, 0700); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var err error

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
		err = daemon()
	case "daemon":
		err = daemon()
	default:
		fmt.Printf("%s not found\n", cmd)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
