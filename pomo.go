package main

import (
	"flag"
	"fmt"
)

func list() {
	fmt.Println("do list")
}

func start() {
	fmt.Println("do start")
}

func stop() {}

func main() {
	flag.Parse()
	args := flag.Args()
	cmd := "list"
	if len(args) != 0 {
		cmd = args[0]
		args = args[1:]
	}
	switch cmd {
	case "list":
		list()
	case "start":
		start()
	case "stop":
		stop()
	default:
		fmt.Printf("%s not found\n", cmd)
	}
}
