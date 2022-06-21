package main

import (
	"log"

	"github.com/onqlavelabs/onqlave.core/internal/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		log.Fatal(err)
	}
}
