package main

import "fmt"

func main() {
	fmt.Println("Welcome to Know Array in GO")
	var programmingLanguages [4]string
	programmingLanguages[0] = "Javascript"
	programmingLanguages[1] = "java"
	programmingLanguages[2] = "python"
	fmt.Println("length is: ", len(programmingLanguages))
	fmt.Println("length is: ", (programmingLanguages))
}

// this normal array wont be used much in Go
