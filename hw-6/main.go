package main

import (
	"flag"
	"log"

	"github.com/speedcuber73/otus-go/hw-6/gocopy"
)

var (
	from   string
	to     string
	offset int64
	limit  int64
)

func init() {
	flag.StringVar(&from, "from", "", "file to read from")
	flag.StringVar(&to, "to", "", "file to read from")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
	flag.Int64Var(&limit, "limit", 0, "max bytes to copy")
}

func main() {
	flag.Parse()

	err := gocopy.Gocopy(from, to, offset, limit)
	if err != nil {
		log.Fatal(err)
	}
}
