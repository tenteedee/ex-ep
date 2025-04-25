package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func runTwoSum() {
	slice := inputSlice()

	fmt.Println("Slice:")
	fmt.Println(slice)

	fmt.Print("Input target: ")
	var target int
	fmt.Scan(&target)

	fmt.Println(twoSum(slice, target))
}
