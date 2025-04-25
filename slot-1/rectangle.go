package main

import "fmt"

func rectangle() {
	fmt.Print("Input width: ")
	var width int
	fmt.Scan(&width)

	fmt.Print("Input height: ")
	var height int
	fmt.Scan(&height)

	fmt.Printf("Area: %d\n", calculateArea(width, height))
	fmt.Printf("Perimeter: %d\n", calculatePerimeter(width, height))
}

func calculateArea(width int, height int) int {
	return width * height
}

func calculatePerimeter(width int, height int) int {
	return 2 * (width + height)
}
