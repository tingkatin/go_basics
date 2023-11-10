package main

import (
	"fmt"
	"greeting"
	"os"
)

func main() {
	fmt.Println(greeting.Say(os.Args[1:]))
}
