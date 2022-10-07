package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("...............")

	ch1, err := read("file1.csv")
	if err != nil {
		panic(fmt.Errorf("could not read file"))
	}

	ch2, err := read("file2.csv")
	if err != nil {
		panic(fmt.Errorf("could not read file"))
	}

	exit := make(chan struct{})

	chM := merge2(ch1, ch2)

	go func() {
		for v := range chM {
			fmt.Println("merge assets ", v)
		}
		close(exit)
	}()

	<-exit
	fmt.Println("All completed, exiting")

}

func merge2(cs ...<-chan []string) <-chan []string {

	chans := len(cs)

	wait := make(chan struct{}, chans)

	out := make(chan []string)

	send := func(c <-chan []string) {
		defer func() {
			wait <- struct{}{}
		}()

		for n := range c {
			out <- n
		}
	}

	for _, c := range cs {
		go send(c)
	}

	go func() {
		for range wait {
			chans--
			if chans == 0 {
				break
			}
		}

		close(out)
	}()

	return out

}

func read(file string) (<-chan []string, error) {
	f, err := os.Open(file)

	if err != nil {
		return nil, fmt.Errorf("opening file %v", err)
	}

	ch := make(chan []string)

	cr := csv.NewReader(f)

	go func() {
		for {
			record, err := cr.Read()
			if err == io.EOF {
				close(ch)
				return
			}
			ch <- record
		}

	}()

	return ch, nil

}
