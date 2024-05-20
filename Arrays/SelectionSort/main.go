package main

import "fmt"

func main() {
	arr := []int{2, 44, 56, 1, 35, 77, 96, 45, 63}
	fmt.Println("Sorted array :", selectionSort(arr))
}

func selectionSort(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[minIndex] > nums[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			nums[i], nums[minIndex] = nums[minIndex], nums[i]
		}
	}
	return nums
}
