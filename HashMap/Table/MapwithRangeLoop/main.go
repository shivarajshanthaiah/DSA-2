package main

import "fmt"

func main() {
	menu := map[string]float32{
		"Soup":   170.50,
		"Salad":  320,
		"Sweets": 488.73,
		"Curry":  100,
	}

	for k, v := range menu {
		fmt.Printf("%s: %f\n", k, v)
	}
}
