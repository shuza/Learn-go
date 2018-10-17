package main

import (
	"os"
	"strconv"
	"fmt"
)

func main() {
	arguments := os.Args
	exitNumber := -1

	if len(arguments) > 0 {
		if cmd, err := strconv.Atoi(arguments[0]); err != nil {
			exitNumber = cmd
		}
	}

	var number int
	for {
		fmt.Print("-> ")
		_, err := fmt.Scanf("%d", &number)
		if err != nil || number == exitNumber {
			break
		}
		fmt.Println("You have entered ", number)
	}
}
