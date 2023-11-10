package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var lc, cc, wc int

	for _, filename := range os.Args[1:] {

		var line_count, character_count, word_count int

		file, err := os.Open(filename)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		scan := bufio.NewScanner(file)
		for scan.Scan() {
			str := scan.Text()

			word_count += len(strings.Fields(str))
			wc += len(strings.Fields(str))
			character_count += len(str)
			cc += len(str)
			line_count++
			lc++
		}

		fmt.Printf(" %9d %9d %9d %s \n", line_count, character_count, word_count, filename)

		file.Close()

	}

	fmt.Printf(" %9d %9d %9d %s \n", lc, cc, wc, "Total")

}
