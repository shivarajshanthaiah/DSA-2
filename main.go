package main

import "fmt"

func main() {
	arr := []int{1, 2, 33, 2, 321, 4, 35, 6, 66, 5, 43, 34, 34, 67, 28}
	fmt.Println(mergeSort(arr))
	fmt.Println(quickSort(arr))
}

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	mid := len(arr) / 2
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

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	left, mid, right := []int{}, []int{}, []int{}

	for _, val := range arr {
		if val < pivot {
			left = append(left, val)
		} else if val == pivot {
			mid = append(mid, val)
		} else {
			right = append(right, val)
		}
	}
	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, mid...), right...)
}
