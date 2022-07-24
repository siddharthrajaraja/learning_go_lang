package main

import "fmt"

func IterateOnMap(m map[string]string) {
	fmt.Println("\nITERATING ON MAP")
	for key, value := range m {
		fmt.Println(key, value)
	}
}

func WaysToCreateMap() {
	// WAY 1
	colors := map[string]string{
		"red":   "RED",
		"green": "GREEN",
	}
	fmt.Printf("%+v", colors)
	IterateOnMap(colors)
	// WAY 2
	names := make(map[string]string)
	names["sid"] = "Siddharth"
	names["raja"] = "Raja"
	fmt.Printf("\n%+v", names)
	delete(names, "sid")
	fmt.Printf("\n%+v", names)

}

func main() {
	WaysToCreateMap()
}
