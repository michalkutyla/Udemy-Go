package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	cards.shuffle().print()
}

func newCard() string {
	return "Five of Diamonds"
}
