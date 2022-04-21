package main

import (
	"flag"
	"fmt"
)

func main() {

	flag.Parse()

	action := flag.Arg(0)

	switch action {

	case "new":
		new()
		break

	case "help":
		helpMessage()
	default:
		helpMessage()
	}
}

func helpMessage() {
	fmt.Println(`Usage: leopard <command> [options]
Commands:
  new		Create a new project
  help		Show this help message`)
}
