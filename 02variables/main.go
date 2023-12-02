package main

import "fmt"

// constant variable
const LogginToken = "This is a Publoc LogginToken because of L is a capital letter"

func main() {
	var username string = "Balaji Narayana"
	fmt.Println(username)
	fmt.Printf("Variable is type of: %T", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)

	var myNumber int = 255
	fmt.Println(myNumber)

	var smallFloat float32 = 255.33455335644
	// decimal precesioj will be less in comparison with float64
	fmt.Println(smallFloat)

	// no var style--this is totally fine, this is allowed only inside methods, not allowed outside of methods
	numberOfUsers := 300000
	fmt.Println(numberOfUsers)
}
