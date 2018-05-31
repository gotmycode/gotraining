package main

import "fmt"

func main() {
	//slice of type int
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//for loop
	for _, num := range numbers {

		//if the
		if num%2 == 0 {
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}

}
