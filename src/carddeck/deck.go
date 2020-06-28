package carddeck

import "fmt"
import "io/ioutil"
import "strings"
import "os"
import "math/rand"
import "time"

type deck []string

func generateNewDeck() deck {
	deck := deck{};
	numbers := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"};
	symbols := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	for _, symbol :=range symbols {
		for _, number := range numbers {
			cardName := number + " of " + symbol;
			deck = append(deck, cardName)
		}
	}
	return deck;
}

func (d deck) print() {
	for index, card:= range d {
		fmt.Println(index, card);
	}
}

func deal(handSize int, d deck) (deck, deck){
	return d[0:handSize], d[handSize:]
}

func writeToFile(d deck, fileName string) {
	deckInString := d.toString();
	ioutil.WriteFile(fileName, []byte(deckInString), 0644);
}

func readFromFile(fileName string) deck {
	deckInBytes, err := ioutil.ReadFile(fileName);
	if (err != nil) {
		if err != nil {
			fmt.Println("File reading error", err)
			os.Exit(1);
		}
	}
	deckInString := string(deckInBytes);
	return fromString(deckInString);
}

func (d deck) toString() string {
	dStrings := []string(d);
	return strings.Join(dStrings, ",");
}

func fromString(deckInString string) deck {
	dStrings := strings.Split(deckInString, ",");
	return deck(dStrings);
}

func (d deck) shuffle() deck {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source);
	for index := range d {
		randomIndex := r.Intn(len(d));
		d[index], d[randomIndex] = d[randomIndex], d[index];
	}
	return d;
}