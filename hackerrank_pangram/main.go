package main

import (
	"fmt"
	"strings"
)

func is_pangram(s string) string {

	collector := make(map[string]int32)
	output := "pangram"

	lowered := strings.ToLower(s)
	splitted := strings.Split(lowered, "")

	for _, letter := range splitted {

		if letter == " " {
			continue
		}

		collector[letter]++

	}

	if len(collector) < 26 {

		output = "not pangram"

	} else {

		for _, value := range collector {

			if value == 0 {
				output = "not pangram"
			}

		}
	}

	return output
}

func main() {

	s := ""
	output := is_pangram(s)

	fmt.Print(output)

}
