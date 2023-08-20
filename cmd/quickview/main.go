package main

import (
	"log"
	"os"

	"github.com/Evertras/quickview/cmd/quickview/cmd"
)

func main() {
	// Log normally tries to go to stderr, switch to stdout for simplicity
	log.SetOutput(os.Stdout)

	if err := cmd.Execute(); err != nil {
		log.Println(err)
	}
}
