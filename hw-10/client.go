package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	flag "github.com/spf13/pflag"
)

func writeRoutine(ctx context.Context, conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if !scanner.Scan() {
				return
			}
			conn.Write([]byte(fmt.Sprintf("%s\n", scanner.Text())))
		}
	}
}

func readRoutine(cancel context.CancelFunc, conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		log.Println(scanner.Text())
	}

	cancel()
}

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

	dialer := &net.Dialer{}
	ctxConn, cancelConn := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancelConn()

	conn, err := dialer.DialContext(ctxConn, "tcp", host+":"+port)
	if err != nil {
		log.Fatalf("Cannot connect: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		readRoutine(cancel, conn)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		writeRoutine(ctx, conn)
		conn.Close()
		wg.Done()
	}()

	wg.Wait()
}
