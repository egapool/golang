package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	// const (
	// 	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	// 	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	// 	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	// 	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	// 	// The remaining values may be or'ed in to control behavior.
	// 	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	// 	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	// 	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
	// 	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	// 	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
	// )
	filepath := filepath.Join(dir, "sample.csv")
	fmt.Println(filepath)
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = file.Truncate(0)
	if err != nil {
		log.Fatal(err)
	}

	// @see https://godoc.org/encoding/csv
	writer := csv.NewWriter(file)
	writer.Write([]string{"Alice", "20"})
	writer.Write([]string{"Bob", "21"})
	writer.Write([]string{"Carol", "22"})
	writer.Flush()

	// reading
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		fmt.Println(record)
	}
}
