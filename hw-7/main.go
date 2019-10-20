package main

import (
	"fmt"
	"log"
	"os"

	"github.com/speedcuber73/otus-go/hw-7/envdir"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("not enough args")
	}
	envPath := os.Args[1]
	command := os.Args[2:]

	out, err := envdir.Run(envPath, command)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(out)
}
