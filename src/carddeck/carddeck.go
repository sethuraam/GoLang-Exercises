package carddeck

func DeckOperations() deck {
	 cardsDeck:=generateNewDeck();
	 cardsDeck = cardsDeck.shuffle()
	 handCards, remainingCards := deal(5, cardsDeck);
	 writeToFile(handCards, "/tmp/cardsInHand")
	 writeToFile(remainingCards, "/tmp/remainingCards")
	 readFromFileDeck := readFromFile("/tmp/remainingCards");
	 readFromFileDeck.print();
	 return cardsDeck;
}