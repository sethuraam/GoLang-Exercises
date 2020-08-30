package main

import "fmt"

type person struct {
	firstName string
	lastName string
	contact contactInfo
}


type contactInfo struct {
	mobileNumber string
	emailAddress string
}

func main() {
	var alex person;
	alex = person{
		firstName: "Alex",
		lastName: "Hunter", 
		contact: contactInfo{
			mobileNumber: "9999999999", 
			emailAddress: "sethuraaam@gmail.com",
			},
		} ;
	alex.print();
	alex.updateName("Jim", "Mcgill");
	alex.print();
}

func (p person) print() {
	fmt.Printf("%+v", p);
}

func (pointerToPerson *person) updateName(firstName string, lastName string) {
	(*pointerToPerson).firstName = firstName;
	(*pointerToPerson).lastName = lastName;
}

func (c contactInfo) print() {
	fmt.Printf("%+v", c);
}