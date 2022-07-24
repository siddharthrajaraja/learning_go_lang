package main

import "fmt"

type Person struct {
	firstName   string
	lastName    string
	contactInfo ContactInfo
}

type ContactInfo struct {
	emaild  string
	zipCode int
}

func (p Person) printPerson() {
	fmt.Printf("\n%+v\n", p)
}

func (p Person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func WaysToCreateStruct() {
	// Way 1 : Not recommended , if fields in druct definition swapped
	siddharth := Person{"Siddharth", "raja", ContactInfo{"siddharthraja9849@gmail.com", 123}}
	fmt.Println(siddharth)

	// Way 2
	siddhu := Person{firstName: "Siddhu", lastName: "raja", contactInfo: ContactInfo{
		emaild: "siddhu@gmail.com", zipCode: 123,
	}}
	fmt.Println(siddhu)

	// Way3
	var monu Person
	fmt.Println(monu) // Zero value assignments
	monu.firstName = "Monu"
	monu.lastName = "Raja"
	monu.contactInfo = ContactInfo{
		emaild:  "siddhu@gmail.com",
		zipCode: 123,
	}
	fmt.Println(monu)
	fmt.Printf("%+v", monu)
	monu.printPerson()
	monu.updateName("Tonu")
	monu.printPerson()
}

func main() {
	WaysToCreateStruct()
}
