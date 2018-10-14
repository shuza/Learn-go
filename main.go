package main

import (
	"fmt"
	"os"
	"strings"
)

/**
 * :=  created by:  Shuza
 * :=  create date:  28-Aug-2018
 * :=  (C) CopyRight Shuza
 * :=  www.shuza.ninja
 * :=  shuza.sa@gmail.com
 * :=  Fun  :  Coffee  :  Code
 **/

func main() {
	arguments := os.Args
	minusI := false

	for i := 0; i < len(arguments); i++ {
		if strings.Compare(arguments[i], "-i") == 0 {
			minusI = true
			break
		}
	}

	if minusI {
		fmt.Println("Got -i parameter!")
		fmt.Print("y/n: ")

		var answer string
		fmt.Scanln(&answer)
		fmt.Println("You have entered: ", answer)

	} else {
		fmt.Println("The -i parameter is not set")
	}
}
