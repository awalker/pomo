package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

const END_ENTRY rune = 0
const FIELD_SEP rune = 0x1f

func rofi(prefix string, field string, msg string) string {
	return fmt.Sprintf("%s%c%s%c%s", prefix, END_ENTRY, field, FIELD_SEP, msg)
}

func rofiMessage(msg string) {
	fmt.Println(rofi("", "message", msg))
}

func list() error {
	rofiMessage("Pomo - pomodoro timers")
	fmt.Println("create")
	fmt.Println("pause")
	fmt.Println("start")
	fmt.Println("stop")
	return nil
}

func start() error {
	fmt.Println("do start")
	return nil
}

func stop() error {
	return nil
}

func create() error {
	// fmt.Println(rofi("", "prompt", "Label"))
	// rofiMessage("Create a new label")
	// fmt.Println("create")
	label, err := rofiInput("Enter Label > ")
	if err != nil {
		return err
	}
	fmt.Println(label)
	return nil
}

func rofiCmd(params ...string) (*exec.Cmd, error) {
	rofiBin, err := exec.LookPath("rofi")
	if err != nil {
		return nil, fmt.Errorf("Could not find rofi command (%w)", err)
	}
	cmd := exec.Command(rofiBin, params...)
	return cmd, nil
}

func rofiMode() error {
	binary := os.Args[0]
	// cmd, err := rofiCmd("-modi", "pomo:"+binary, "-show", "pomo")
	cmd, err := rofiCmd("-prompt", "pomo> ", "-dmenu")
	if err != nil {
		return err
	}
	var out bytes.Buffer

	cmd.Stdout = &out
	// cmd.Stdin =
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("Could not run rofi command (%w)", err)
	}
	fmt.Println(out.String())
	return nil
}

func rofiInput(prompt string) (string, error) {
	// bla=$(rofi -dmenu -input /dev/null -p "Enter Text > ")
	cmd, err := rofiCmd("-dmenu", "-input", "/dev/null", "-p", prompt)
	if err != nil {
		return "", err
	}
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("Could not get input from rofi command (%w)", err)
	}
	return out.String(), nil
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
		err = rofiMode()
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
