package main

import "fmt"

func main() {

	s := 12345
	a := 'a'
	fmt.Println(a)

	for i, v := range s {
		fmt.Println(i, v)
	}
}
