package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Not enough arguments")
	}
	var text []string
	old, new := os.Args[1], os.Args[2]
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		split := strings.Split(scan.Text(), old)
		textPerLine := strings.Join(split, new) + "\n"

		text = append(text, textPerLine)
	}

	// Create New TXT File
	file, fileErr := os.Create("result.txt")
	if fileErr != nil {
		fmt.Println("An error occurred: ", fileErr)
		return
	}
	defer file.Close()

	// Write each line to the file on a new line
	for _, line := range text {
		_, writeErr := file.WriteString(line)
		if writeErr != nil {
			fmt.Println("An error occurred: ", writeErr)
			return
		}
	}
}
