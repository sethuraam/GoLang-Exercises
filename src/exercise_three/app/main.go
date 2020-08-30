package main

import "fmt"

func main() {
	var color map[string]string;
	color = map[string]string{
		"red" : "FF0000",
		"green" : "00FF00",
		"blue" : "0000FF",
	};
	delete(color, "red")
	printMap(color);
}

func printMap(mapToBePrinted map[string]string) {
	for key,value:=range(mapToBePrinted) {
		fmt.Println(key +  ": " + value);
	}
}