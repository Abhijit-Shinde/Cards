package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, cardSuite := range cardSuites {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+cardSuite)
		}
	}

	return cards
}

func deal(cards deck, upto int) (deck, deck) {
	return cards[0:upto], cards[upto:]
}

// Receiver on a function
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) stringToByte() string {
	return strings.Join(d, ", ")
}

func (d deck) saveToFile(file string) error {
	return os.WriteFile(file, []byte(d.stringToByte()), 0666)
}

func deckFromFile(file string) deck {
	cardsFromFile, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Split(string(cardsFromFile), ",")
}

func (d deck) shuffle() {
	// for originalPos := range d {
	// 	newPos := rand.Intn(originalPos + 1)
	// 	d[originalPos], d[newPos] = d[newPos], d[originalPos]
	// }

	rand.Shuffle(len(d),func (i int,j int)  {
		d[i], d[j] = d[j], d[i]
	})
}
