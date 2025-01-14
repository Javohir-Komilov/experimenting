package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type MyReader struct {
	data string
}

func (r *MyReader) Read(p []byte) (n int, err error) {
	if len(r.data) == 0 {
		return 0, io.EOF
	}
	n = copy(p, r.data)
	r.data = r.data[n:]
	return n, nil
}

func main() {
	reader := &MyReader{data: "Hello, Golang!"}
	buf := make([]byte, 8)

	for {
		n, err := reader.Read(buf)
		fmt.Print(string(buf[:n]))
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			break
		}
	}

	fmt.Println()

	strReader := strings.NewReader("Go is awesome!")
	io.Copy(os.Stdout, strReader)
}
