package main

import "fmt"

func main() {
	// colors := map[string]string{
	// 	"black": "#00000",
	// 	"white": "#ffffff",
	// 	"red":   "#ff0000",
	// }

	colors := make(map[string]string)

	colors["red"] = "#ff0000"

	fmt.Println(colors)
	printMap(colors)
	delete(colors, "yellow ")
}

func printMap(c map[string]string) {
	for color, value := range c {
		fmt.Println("Hex code for " + color + " is " + value)
	}
}
