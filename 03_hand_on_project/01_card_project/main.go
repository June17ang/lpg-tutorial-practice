package main

func main() {
	cards := newDeckFromFile("my_cards")
	cards.printOutResult()

	cards.shuffle()
	cards.printOutResult()
}
