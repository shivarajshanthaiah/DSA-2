package main

import "fmt"

func main() {
	arr := []int{3, 34, 63, 73, 23, 56, 34, 15, 80}
	fmt.Println("Sorted Array :", insertionSort(arr))
}

func insertionSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		temp := nums[i]
		j := i - 1
		for j >= 0 && temp < nums[j] {
			nums[j+1], nums[j] = nums[j], temp
			j--
		}
	}
	return nums
}
