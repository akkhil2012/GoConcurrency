package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("In main.go class")

	records, err := readCSV("testFile.csv")

	if err != nil {
		log.Fatal("Could not read csv file")
	}

	//for val := range titleize(records) {
	for val := range records {
		fmt.Println("--> ", val)
	}

	fmt.Println(" records are : ", records)
}

func titleize(strC <-chan []string) <-chan []string {
	ch := make(chan []string)

	go func() {
		for val := range strC {
			val[0] = strings.Title(val[0])
			val[1], val[2] = val[2], val[1]

			ch <- val
		}
		close(ch)
	}()

	return ch
}

func readCSV(file string) (<-chan []string, error) {
	f, err := os.Open(file)

	if err != nil {
		return nil, fmt.Errorf("opening file %w", err)
	}
	fmt.Println(" while reading CSV  : ", f)

	ch := make(chan []string)

	go func() {
		cr := csv.NewReader(f)
		cr.FieldsPerRecord = 3

		for {
			record, err := cr.Read()
			if errors.Is(err, io.EOF) {
				close(ch)

				return
			}
			ch <- record

		}
	}()

	return ch, nil
}
