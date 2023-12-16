package main

import "fmt"

func main() {
	// var colors map[string]string

	// colors := make(map[string]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#fffff",
	}

	// colors["white"] = "#fffff"

	// delete(colors, "white")
	// structName.white

	// fmt.Println(colors)
	printMap(colors)
}

// if you do not assign a value, go will initialize it with a zero value

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}

// Map
// all keys must be the same type
// use to represent a collection of related properties
// all values must be the same type
// do not need to know all the keys at compile time
// keys are indexed - we can iterate over them
// reference type

// Struct
// values can be of diff type
// you need to know all the diff fields at compile time
// keys do not support indexing
// use to represent a thing with a lot of diff properties
// value type


