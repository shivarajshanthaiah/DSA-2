package main

import "fmt"

func main() {
	arr := []int{29, 40, 44, 25, 58, 66}
	fLar := arr[0]
	secLar := arr[1]

	if fLar < secLar {
		fLar, secLar = secLar, fLar
	}

	for i := 2; i < len(arr); i++ {
		if arr[i] > fLar {
			secLar, fLar = fLar, arr[i]
		} else if arr[i] > secLar && arr[i] != fLar {
			secLar = arr[i]
		}
	}
	fmt.Println(secLar)

}
