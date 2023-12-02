package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.saveToFile("test.txt")
	cards = deckFromFile("test.txt")
	dealt,remain := deal(cards,2)
	dealt.print()
	remain.print()
	cards.print()
}
