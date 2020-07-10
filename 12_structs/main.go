package main

import (
	"fmt"
	"strconv"
)

// Define person struct

type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hell, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "Männlich" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct

	/* person1 := Person{firstName: "Konstantin", lastName: "Komendantov", city: "Heilbronn", gender: "Männlich", age: 21} */
	person2 := Person{"Konstantin", "Komendantov", "Neckarsulm-Amorbach", "Männlich", 21}
	person3 := Person{"Susanne", "Schmidt", "Neckarsulm-Amorbach", "Weiblich", 21}

	/* fmt.Println(person2)

	fmt.Println(person2.firstName)
	person2.age++
	fmt.Println(person2) */

	person2.hasBirthday()
	fmt.Println(person2.greet())
	person3.getMarried("Frank")
	fmt.Println(person3.greet())

}
