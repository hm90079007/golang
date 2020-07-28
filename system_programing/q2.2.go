package main

import (
	"encoding/csv"
	"os"
)

func main() {
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
}
