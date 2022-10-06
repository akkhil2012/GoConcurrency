package main

import "fmt"

// functions

func sumVar(nums ...int) int {
	total := 0
	for n := range nums {
		total += n
	}
	return total
}

func add(a int, b int) int {
	return a + b
}

func addAll(a, b, c int) int {
	return a + b + c
}

func multiAdd(a, b string) (string, string) {
	front := a + b
	rev := b + a
	return front, rev
}

func main() {

	nums := []int{1, 2, 3, 4}
	var sumAllNum = sumVar(nums...)
	fmt.Println(" Sum is========================== ", sumAllNum)

	a, b := multiAdd("akhil", "gupta")
	fmt.Println("front ", a)
	fmt.Println("reverse ", b)

	add := add(3, 4)
	fmt.Println("sum is", add)

	addAll := addAll(3, 4, 9)
	fmt.Println("sum All  is", addAll)

	m := make(map[string]int)
	m["akkhil"] = 1
	m["gupta"] = 2
	fmt.Println(">>>", m)
	n := map[string]int{"india": 1, "hyd": 2}
	fmt.Println("modified map ", n)

	// range is used as iterator
	for k, v := range n {
		fmt.Println("%s -> %d\n", k, v)
	}

}
