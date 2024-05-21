package main

import "fmt"

func main() {
	arr := []int{23, 5, 76, 3, 86, 53, 54, 6, 3}
	fmt.Println(quickSort(arr))
}

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	pivot := nums[len(nums)/2]
	left, right := []int{}, []int{}

	for _, n := range nums {
		if n < pivot {
			left = append(left, n)
		} else if n > pivot {
			right = append(right, n)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}
