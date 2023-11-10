package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	words := make(map[string]int)
	scan := bufio.NewScanner(os.Stdin)

	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		word := scan.Text()

		words[word]++
	}

	// Create New TXT File
	file, fileErr := os.Create("result.txt")
	if fileErr != nil {
		fmt.Println("An error occurred: ", fileErr)
		return
	}
	defer file.Close()

	// Write each line to the file on a new line
	for word, count := range words {
		_, writeErr := file.WriteString(word + ": " + fmt.Sprint(count) + "\n")
		if writeErr != nil {
			fmt.Println("An error occurred: ", writeErr)
			return
		}
	}

}
