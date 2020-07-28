package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("multiwriter.txt")
	if err != nil {
		panic(err)
	}
	writer := io.MultiWriter(file, os.Stdout)
	//	io.WriteString(writer, "io.MultiWriter example\n")
	fmt.Fprintf(writer, "hello %s\n", "enako")
}
