package main

import (
	"fmt"
)

func main() {

	arr := [6]int{1, 3, 5, 6, 7, 8}

	for v := 0; v < len(arr); v++ {
		if v%2 == 0 {
			fmt.Println("--> ", arr[v], " index is ", v)
		}
	}

	if num := 10; num < 0 {
		fmt.Println("its negative..")
	} else if num < 10 {
		fmt.Println("its one digit ..")
	} else {
		fmt.Println("has  multiple....")
	}

	for i := 1; i < 5; i++ {
		fmt.Println("i: ", i)
	}

	j := 1
	for j <= 3 {
		fmt.Println("j: ", j)
		j = j + 1
	}

	for {
		fmt.Println("loop ")
		break
	}

	var score int = 100

	marks := 33

	var c, d int = 3, 4

	const limit = 222

	fmt.Println("constant is ", int64(limit))
	fmt.Println("score is ", score)
	fmt.Println("marks  is ", marks)

	fmt.Println("c  is ", c)
	fmt.Println("d  is ", d)
	fmt.Println("test Class")
	fmt.Println("1+1 is ", 1+1)
	fmt.Println("OR is ", false || true)
	fmt.Println("AND is ", false && true)
}
