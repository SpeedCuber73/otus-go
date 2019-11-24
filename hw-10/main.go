package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	flag "github.com/spf13/pflag"
)

var timeoutStr string

func init() {
	flag.StringVar(&timeoutStr, "timeout", "10s", "timeout for connection")
}

func main() {
	flag.Parse()

	if !strings.HasSuffix(timeoutStr, "s") {
		log.Fatal("bad timeout unit")
	}

	timeout, err := strconv.Atoi(timeoutStr[:len(timeoutStr)-1])
	if err != nil {
		log.Fatal("bad timeout number value")
	}

	if len(flag.Args()) != 2 {
		log.Fatal("bad parameters")
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	fmt.Println(timeout, host, port)
}
