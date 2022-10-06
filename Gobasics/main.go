package main

import (
	"errors"
	"fmt"
)

func testError(arg int) (int, error) {
	if arg == 10 {
		return -1, errors.New("11 is a restricion")
	}
	return arg + 3, nil
}

func copyByValue(iVal int) {
	iVal = 0
}

func copyByPointe(iPtr *int) {
	*iPtr = 0
}

/*
interfcae
*/

type geometry interface {
	area() float64
	perim() float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Area of geometry:", g.area())

}

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 33
	return &p
}

type rect struct {
	width, height float64
}

func main() {

	_, answer := testError(10)
	if answer != nil {
		fmt.Println(" res is>>>>>>>>>>>>>> : ", answer)
	} else {
		fmt.Println(" error ...............")
	}

	r := rect{width: 10, height: 5}
	fmt.Println("area---------------> ", r.area())

	fmt.Println("person: ", person{"akkhl", 30})
	//fmt.Println("person: ", )

	pp := person{"gupta", 44}

	fmt.Println(" first ", pp)
	pp.age = 55
	fmt.Println(" sec ", pp)

	//fmt.Println("")

	fmt.Println("<--------------------------------->")

	i := 100

	fmt.Println("1:1 ", i)

	copyByValue(i)
	fmt.Println("1:3: ", i)

	fmt.Println("2:1 ", i)

	copyByPointe(&i)

	fmt.Println("2:3 ", i)

}
