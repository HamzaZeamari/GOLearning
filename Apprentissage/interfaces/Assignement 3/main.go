package main

import (
	"fmt"
	"os"
)

type logWriter struct {
}

func main() {
	File := os.Args[1]

	f, err := os.Open(File)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log := logWriter{}
	bs := make([]byte, 500)
	f.Read(bs)
	log.Write(bs)
	f.Close()
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println(len(bs))
	return len(bs), nil
}
