package main

import "fmt"

func sliceAggregate() {
	slice := inputSlice()

	sum, min, max, avg := aggregate(slice)

	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Min: %d\n", min)
	fmt.Printf("Max: %d\n", max)
	fmt.Printf("Avg: %.2f\n", avg)

	fmt.Println("Sorted slice:")
	fmt.Println(sort(slice))
}

func inputSlice() []int {
	fmt.Print("Input size of slice: ")
	var size int
	fmt.Scan(&size)

	slice := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Printf("Input num %d: ", i+1)
		fmt.Scan(&slice[i])
	}

	return slice
}

func aggregate(slice []int) (int, int, int, float64) {
	sum := 0
	min := slice[0]
	max := slice[0]

	for _, value := range slice {
		sum += value

		if value < min {
			min = value
		}

		if value > max {
			max = value
		}
	}

	avg := float64(sum) / float64(len(slice))

	return sum, min, max, avg
}

func sort(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}

	return slice
}
