package main

import "fmt"

func main() {

	name := map[int]string{
		1: "Shivaraj",
		2: "Eaby",
		3: "Binal",
		4: "Viswanath",
		5: "Sreesankar",
	}

	fmt.Println(name)
	key := 3
	fmt.Printf("%d: %s\n", key, name[key])
}
