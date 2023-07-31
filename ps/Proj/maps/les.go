/*
This has all he actual lessons work
*/

package main

import (
	f "fmt"
)

func method1() {

	color := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}

	hea("The following is just printing out the map which was just declared")
	f.Println(color)
}

func method2() {

	var colors map[string]string

	// Note if you dont declare any variable it will be an empty map

	hea("Second Method usign var")
	f.Println(colors)
}

func method3() {

	colors := make(map[string]string)

	// Defining the values method
	colors["white"] = "#FFFFFF"

	hea("Third Method using make constructor")
	f.Println(colors)
}

func method4() {

	colors := make(map[int]string)

	// Defining the values method
	colors[100] = "#FFFFFF"

	hea("Using int in the map declaration")
	f.Println(colors)
}

func method5() {

	colors := make(map[int]string)

	// Defining the values method
	colors[100] = "#FFFFFF"

	delete(colors, 100)

	hea("Deleting Keys in the map declaration")
	f.Println(colors)
}

// ++++++++++++++++++++ GROUP +++++++++++++++++++++++++++


/*
Maps Iteration
*/

func method6() {

	color := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
		"white": "#FFFFFF",
	}

	hea("Iterating over a created map, wher the iterating over the map is done in another function")
	f.Println(color)


	m6printMap(color)

}

// The following funtion is for method6

func m6printMap(c map[string]string) {

	for color, hex := range c {
		f.Println("Hex Code for ", color, "is suck her" ,hex)
	}
}


// ++++++++++++++++++++ GROUP +++++++++++++++++++++++++++

