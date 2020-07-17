package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
)

func main() {
	// this is where stuff happens

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	text, _ := reader.ReadString('\n')
	fmt.Printf("Hello %v\n", text)
	fmt.Printf("We are using Go %v running in %v\n", runtime.Version(), runtime.GOOS)
}
