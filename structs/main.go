package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func WaysToCreateStruct() {
	// Way 1 : Not recommended , if fields in druct definition swapped
	siddharth := Person{"Siddharth", "raja"}
	fmt.Println(siddharth)

	// Way 2
	siddhu := Person{firstName: "Siddhu", lastName: "raja"}
	fmt.Println(siddhu)

	// Way3
	var monu Person
	fmt.Println(monu) // Zero value assignments
	monu.firstName = "Monu"
	monu.lastName = "Raja"
	fmt.Println(monu)
	fmt.Printf("%+v", monu)

}

func main() {
	WaysToCreateStruct()
}
