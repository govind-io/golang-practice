package main

import "fmt"

// func main() {
// 	var randomNums []int=

// 	fmt.Println(randomNums)

// }

func main() {

	lists := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index, value := range lists {
		fmt.Println(value, index)
	}

}
