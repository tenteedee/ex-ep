package main

import "fmt"

func checkString() {
	fmt.Print("Input string: ")
	var str string
	fmt.Scan(&str)

	fmt.Println(len(str)%2 == 0)
}
