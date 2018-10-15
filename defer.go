package main

import "fmt"

func a1() {
	for i := 0; i < 3; i++ {
		defer fmt.Print(i, " ")
	}
}

func a2() {
	for i := 0; i < 3; i++ {
		defer func() { fmt.Print(i, " ") }()
	}
}

func a3() {
	for i := 0; i < 3; i++ {
		defer func(n int) { fmt.Print(n, " ") }(i)
	}
}

func try(x, y int) int {
	r := x / y
	fmt.Println("try is going to return : ", r)
	return r
}

func main() {
	a1()
	fmt.Println()
	a2()
	fmt.Println()
	a3()
	fmt.Println()

	r := try(4, 2)
	fmt.Println("4 / 2 = ", r)
	defer fmt.Println("error at 4 / 2")

	r = try(4, 0)
	fmt.Println("4 / 0 = ", r)
	defer fmt.Println("error at 4 / 0")

	r = try(5, 2)
	fmt.Println("5 / 0 = ", r)
	defer fmt.Println("error at 5 / 0")
}
