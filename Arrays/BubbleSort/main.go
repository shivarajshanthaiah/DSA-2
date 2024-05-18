package main

import "fmt"

func main() {
	arr := []int{34, 23, 45, 89, 23, 1, 3}
	fmt.Println("Sorted array :", bubbleSort(arr[:]))
	fmt.Println("Sorted array :", optimizesBubbleSort(arr[:]))

}

// descending
// time complexity => O(n^2)
func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] < nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

// ascending and this is just optimised sorting method for the above algorithm which reduces time complexity
// time complexity => O(n)
func optimizesBubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		swap := false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
	return nums
}
