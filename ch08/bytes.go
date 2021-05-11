package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("This is"))
	fmt.Fprintf(&buffer, " a string!\n") // write to buffer
	buffer.WriteTo(os.Stdout) //
	buffer.WriteTo(os.Stdout) // print nothing

	buffer.Reset()
	buffer.Write([]byte("Mastering Go!"))
	fmt.Println(buffer.String())


	r := bytes.NewReader([]byte(buffer.String()))

	for {
		b := make([]byte, 3)
		n, err := r.Read(b)
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Read %s Bytes: %d\n", b, n)
	}
}
