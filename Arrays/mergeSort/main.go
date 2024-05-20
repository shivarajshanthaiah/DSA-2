package main

import "fmt"

func main() {
	arr := []int{47, 23, 13, 5, 6, 82, 55, 62, 56}
	sortedArray := mergeSort(arr)
	fmt.Println(sortedArray)
}

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	mid := (len(arr) / 2)
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	combined := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			combined = append(combined, left[i])
			i++
		} else {
			combined = append(combined, right[j])
			j++
		}
	}

	if i < len(left) {
		combined = append(combined, left[i:]...)
	}
	if j < len(right) {
		combined = append(combined, right[j:]...)
	}
	return combined
}
