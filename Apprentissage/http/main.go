package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {
	resp, err := http.Get("http://www.google.com")
	if err == nil {
		fmt.Println(resp)
	} else {
		fmt.Println("Error")
		os.Exit(1)
	}

	/*bs := make([]byte, 9999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))*/

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Ecrit le nombre de bytes : ", len(bs))
	return len(bs), nil
}
