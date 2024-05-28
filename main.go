package main

import "fmt"

func main() {
	arr := []int{1, 21, 33, 4, 50, 6, 7}
	fmt.Println(mergeSort(arr))
	fmt.Println(quickSort(arr))
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

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[len(arr)/2]
	left, right := []int{}, []int{}

	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		}
	}

	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, pivot), right...)
}
