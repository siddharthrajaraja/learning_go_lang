package main

import "fmt"

func newCard() string {
	return "5 of diamonds"
}

func main() {
	//cards := newDeck()
	//cards.PrintDeck()
	//firstHalf, secondHalf := Deal(cards, 3)
	//fmt.Println(firstHalf)
	//fmt.Println(secondHalf)
	//cards.saveToFile("sample.csv")
	data := newDeckFromFile("sample.csv")
	fmt.Println("Read data", data)
	data.PrintDeck()
	data.shuffle()
	data.PrintDeck()

}
