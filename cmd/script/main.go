package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//read N for
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("----------------------------------")
	fmt.Print("Input N: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
}
