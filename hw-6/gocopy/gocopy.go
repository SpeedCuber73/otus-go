package gocopy

import (
	"fmt"
	"io"
	"log"
	"os"
)

const copyStep = 10

func limitTo(x int64, max int64) int64 {
	if x < 0 {
		return 0
	}
	if x > max {
		return max
	}
	return x
}

// Gocopy copy
func Gocopy(from, to string, offset, limit int64) error {
	srcFile, err := os.Open(from)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	fileStat, err := srcFile.Stat()
	if err != nil {
		return err
	}

	var size int64
	if limit > 0 {
		size = limitTo(fileStat.Size()-offset, limit)
	} else {
		size = fileStat.Size() - offset
	}

	srcFile.Seek(offset, io.SeekStart)
	srcReader := io.LimitReader(srcFile, size)

	dstFile, err := os.Create(to)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	var copied int64
LOOP:
	for {
		_, err := io.CopyN(dstFile, srcReader, copyStep)
		switch err {
		case nil:
		case io.EOF:
			break LOOP
		default:
			log.Fatal("cannot copy")
		}
		copied += copyStep
		fmt.Println("copied ", copied, "bytes out of ", size)
	}

	return nil
}
