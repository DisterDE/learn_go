package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, c := range d {
		fmt.Println(i, c)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeck() deck {
	cards := deck{}

	cardValues := []string{"Ace", "Two", "Three", "Four"}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}

	var sb strings.Builder
	for _, value := range cardValues {
		for _, suit := range cardSuits {
			sb.WriteString(value)
			sb.WriteString(" of ")
			sb.WriteString(suit)
			cards = append(cards, sb.String())
			sb.Reset()
		}
	}

	return cards
}

func newDeckFromFile(filename string) deck {
	c, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return strings.Split(string(c), ",")
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newI := r.Intn(len(d) - 1)
		d[i], d[newI] = d[newI], d[i]
	}
}
