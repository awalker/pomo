package rofi

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

const END_ENTRY rune = 0
const FIELD_SEP rune = 0x1f

func str(prefix string, field string, msg string) string {
	return fmt.Sprintf("%s%c%s%c%s", prefix, END_ENTRY, field, FIELD_SEP, msg)
}

func Message(msg string) {
	fmt.Println(str("", "message", msg))
}

func Cmd(params ...string) (*exec.Cmd, error) {
	rofiBin, err := exec.LookPath("rofi")
	if err != nil {
		return nil, fmt.Errorf("Could not find rofi command (%w)", err)
	}
	cmd := exec.Command(rofiBin, params...)
	return cmd, nil
}

func Dmenu(prompt string, message string, menuOptions ...string) (string, error) {
	options := make([]string, 5)
	options[0] = "-dmenu"
	options[1] = "-p"
	options[2] = prompt
	if message != "" {
		options = append(options, "-mesg")
		options = append(options, message)
	}
	cmd, err := Cmd(options...)
	if err != nil {
		return "", err
	}
	var out bytes.Buffer
	menu := strings.Join(menuOptions, "\n")

	cmd.Stdout = &out
	cmd.Stdin = strings.NewReader(menu)
	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("Could not run rofi command (%w)", err)
	}
	choice := strings.TrimSpace(out.String())
	return choice, nil
}

func Input(prompt string) (string, error) {
	// bla=$(rofi -dmenu -input /dev/null -p "Enter Text > ")
	cmd, err := Cmd("-dmenu", "-input", "/dev/null", "-p", prompt)
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
