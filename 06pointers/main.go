package main

import "fmt"

func main() {
	fmt.Println("welocme to pointers")
	myNumber := 45
	ptr := &myNumber
	fmt.Println("ptr is:", ptr)
	fmt.Println("ptr is:", *ptr+2)
	// for reference types,while passing to the functions/methods actuall refrencce type variable refrence will be passes and some data types copy will be passed so to handle such kind of irreguralites we use pointer
}
