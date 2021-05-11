package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func printFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	io.WriteString(os.Stdout, scanner.Text())
	io.WriteString(os.Stdout, "\n")


	for scanner.Scan() {
		//time.Sleep(1 * time.Second)
		io.WriteString(os.Stdout, scanner.Text())
		io.WriteString(os.Stdout, "\n")
	}
	return nil
}

func main() {
	filename := ""
	arguments := os.Args
	if len(arguments) == 1 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	for i := 1; i < len(arguments); i++ {
		filename = arguments[i]
		err := printFile(filename)
		if err != nil {
			fmt.Println(err)
		}
	}
}


//➜  ch08 git:(master) ✗ go run cat.go goFind.go
//# command-line-arguments
//./goFind.go:34:6: main redeclared in this block
//previous declaration at ./cat.go:31:6
