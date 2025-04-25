package main

import "fmt"

func main() {
	for {
		fmt.Println("=======================")
		fmt.Println("0. Exit")
		fmt.Println("1. Rectangle")
		fmt.Println("2. Check if string length is even")
		fmt.Println("3. Slice aggregate")
		fmt.Println("4. Two sum")

		fmt.Print("Input choice: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			rectangle()
		case 2:
			checkString()
		case 3:
			sliceAggregate()
		case 4:
			runTwoSum()
		case 0:
			return
		default:
			fmt.Println("Invalid choice!")
			continue
		}

	}
}
