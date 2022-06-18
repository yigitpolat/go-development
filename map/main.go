package main

import "fmt"

func main() {

	// assing with zero value which is ""
	// var colors map[string]string

	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff00000",
		"green": "#00ff00",
	}

	colors["white"] = "#ffffff"

	fmt.Println(colors)

	delete(colors, "red")

	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("hex code for", key, "is", value)
	}
}
