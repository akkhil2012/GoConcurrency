package main

import (
	"fmt"
)

type singleton struct {
}

var singleInstance *singleton

func getInstance() *singleton {
	if singleInstance == nil {
		singleInstance = &singleton{}
		fmt.Println("Created new instance..................")
	} else {
		fmt.Println("Instance already there")
	}

	return singleInstance
}

func main() {

	for i := 0; i < 30; i++ {
		go getInstance()
	}

}
