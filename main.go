package main

import (
	"github.com/devlup-labs/django-dep/cmd"
	"log"
	"os"
)

func main() {
	if err := cmd.Run(os.Args[1:]); err != nil {
		log.Fatalf("unable to run the command %s ", err)
	}
}
