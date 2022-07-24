package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// type of deck which is a slice of strings
type deck []string

func newDeck() deck {
	myDeck := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			myDeck = append(myDeck, cardValue+" of "+cardSuit)
		}
	}

	return myDeck
}

func (d deck) PrintDeck() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func Deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) convertDeckToString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.convertDeckToString()), 0666)
}

func newDeckFromFile(filename string) deck {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}
	message := strings.Split(string(data), ",")
	return deck(message)
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())

	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		t := d[newPosition]
		d[newPosition] = d[i]
		d[i] = t
	}
}
