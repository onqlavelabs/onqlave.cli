package main

import (
	"log"

	"github.com/onqlavelabs/onqlave.core/cmd/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		log.Fatal(err)
	}
}
