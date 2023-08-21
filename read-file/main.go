package main

import (
	"fmt"
	"io"
	"os"
)

type readFile struct{}

func main() {
	f := openFile(os.Args[1])

	// w := readFile{}

	io.Copy(os.Stdout, f)

}

func openFile(filename string) *os.File {
	data, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return data
}

func (readFile) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("We just write: ", len(bs))

	return len(bs), nil
}
