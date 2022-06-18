package main

import "fmt"

func main() {

	ints := []int{1, 2, 3, 4, 5}
	letters := []string{"a", "b", "c", "d", "e"}

	for _, intt := range ints {
		for _, letter := range letters {
			fmt.Println(intt, letter)
		}
	}

}
