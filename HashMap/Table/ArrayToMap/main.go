package main

import "fmt"

func main() {
	arr := [10]int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	ans := arrToMap(arr[:])
	print(ans)
}

func print(sqr map[int]int) {
	fmt.Println(sqr)
}

func arrToMap(num []int) map[int]int {
	sqr := make(map[int]int)

	for _, val := range num {
		sqr[val] = val * val
	}
	return sqr
}
