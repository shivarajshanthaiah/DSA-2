package main

import (
	"fmt"
)

func main() {
	str := "abcdefsdkfwdjhdssldkaswasasiuqdsd"

	dup := make(map[byte]byte)

	for i := 0; i < len(str); i++ {
		dup[str[i]]++
	}

	res := []byte{}
	for key, val := range dup {
		if val == 1 {
			res = append(res, key)
		}
	}

	fmt.Println(string(res))
}
