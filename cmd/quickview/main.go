package main

import (
	"log"

	"github.com/Evertras/quickview/cmd/quickview/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Println(err)
	}
}
