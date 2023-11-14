package main

import (
	"baseframe/internal/goserver"
	"os"
)

func main() {
	command := goserver.NewGoServerCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
