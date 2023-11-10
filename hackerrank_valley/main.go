package main

// QUESTION LINK: https://www.hackerrank.com/challenges/three-month-preparation-kit-counting-valleys/problem

import (
	"fmt"
	"strings"
)

func main() {
	str := "UDDDUDUU"
	var altitude int32
	var counter int32

	splitted := strings.Split(str, "")

	for _, str := range splitted {

		if str == "U" {
			altitude++

			if altitude == 0 {
				counter++
			}

		} else {

			altitude--
		}
	}

	fmt.Print(counter)
}
