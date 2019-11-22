package main

import "fmt"

func testMap() {
	colors := make(map[string]string)

	colors["red"] = "blood"
	colors["blue"] = "sky"

	iterateOverMaps(colors)

	//fmt.Println(colors)

}

func iterateOverMaps(color map[string]string) {
	for name, second := range color {
		fmt.Println("hey", name, "fdsf", second)
	}
}
