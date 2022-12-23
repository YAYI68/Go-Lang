package main

func main() {
	cards := newCard()
	cards.shuffle()
	cards.print()

}

// // cards.saveToFile("cardgame")
// newcard := getDeckFromFile("cardgame")
// newcard.print()
// // fmt.Println(newcard)

// func newCard() string {
// 	return "Five of spades"
// }

// // var card string = "Ace of spades "
// 	// card := newCard()
// 	cards := []string{"Ace of spades", newCard()}
// 	cards = append(cards, "Ace of Diamond")
// 	for i, card := range cards {
// 		fmt.Println(i, card)
// 	}
