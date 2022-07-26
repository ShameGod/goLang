package main

import "fmt"

func main() {
	var x [5]int
	x[1] = 5
	fmt.Println(x[1])
	var y = [...]int{1, 2, 3, 4, 5}

	for i, v := range y {
		fmt.Println("index %d, value %d", i, v)
	}

}
