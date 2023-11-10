package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "SOSSSSSQSSOR"
	var counter int32
	var groups [][]string
	var arr []string

	splitted := strings.Split(s, "")

	for index, letter := range splitted {
		arr = append(arr, letter)
		if (index+1)%3 == 0 {
			groups = append(groups, arr)
			arr = nil
		}
	}

	for _, group := range groups {
		for index, letter := range group {
			if (index == 0 || index == 2) && letter != "S" {
				counter++
			}

			if index == 1 && letter != "O" {
				counter++
			}
		}
	}

	fmt.Print(counter)
}
