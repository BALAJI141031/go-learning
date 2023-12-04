package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time package in Glolang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	todayDate := presentTime.Format("01-02-2006 15:04:05 Monday")
	fmt.Println(todayDate)
}
