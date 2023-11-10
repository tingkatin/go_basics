package main

import (
	"fmt"
	"os"
)

func main() {
	var sum float64
	var avg float64
	var n int

	for {
		var value float64

		if _, err := fmt.Fscanln(os.Stdin, &value); err != nil {
			break
		}

		sum += value
		n++
	}

	if n == 0 {
		fmt.Fprintln(os.Stderr, "No values found")
		os.Exit(-1)
	}

	avg = sum / float64(n)
	fmt.Printf("The average is: %.2f", avg)

}
