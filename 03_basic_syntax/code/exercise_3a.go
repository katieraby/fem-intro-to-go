package main

import (
	"fmt"
)

func main () {
	sentence := "Here is a sentence"

	for index, letter := range sentence {
		if index % 2 != 0 {
			fmt.Println("letter:", string(letter))
		}
	}
}
