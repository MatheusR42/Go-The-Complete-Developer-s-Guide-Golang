package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("my_cards")
	cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
	// hand, remaingCards := deal(cards, 5)
	// fmt.Println(hand)
	// fmt.Println(remaingCards)
}