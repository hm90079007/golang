package main

import (
	"encoding/csv"
	"io"
	"os"
)

func main() {
	//write file
	records := [][]string{
		[]string{"enako", "E"},
		[]string{"nekomu", "I"},
		[]string{"nanase", "A"},
		[]string{"rara", "J"},
	}

	file, err := os.Create("test1.csv")
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(file) //interface

	for _, record := range records {
		if err := writer.Write(record); err != nil {
			panic(err)
		}
	}

	writer.Flush()

	if err := writer.Error(); err != nil {
		panic(err)
	}

	// read file
	file1, err := os.Open("test1.csv")
	if err != nil {
		panic(err)
	}
	defer file1.Close()

	// create file
	file2, err := os.Create("test2.csv")
	if err != nil {
		panic(err)
	}
	defer file2.Close()

	//write file file1 --> file2
	io.Copy(file2, file1)

}
