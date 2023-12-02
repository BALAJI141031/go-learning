package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Here is the example to take user Input")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Rating for my GO learning")
	input, _ := reader.ReadString('\n')
	fmt.Println("your rating is", input)
}
