package main

import "fmt"

func main() {
	arr := []int{}

	for i := 7; i < 100; i++ {
		if i%7 == 0 {
			arr = append(arr, i)
		}
	}
	fmt.Println(arr)

	sqr := make(map[int]int)
	for _, val := range arr {
		sqr[val] = val * val
	}
	fmt.Println(sqr)

}
