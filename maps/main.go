package main

import "fmt"

func main() {
	// Ways to create maps
	// var colors map[string]string

	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
	}

	colors["white"] = "#FFFFFF"

	fmt.Println(colors)

	delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
