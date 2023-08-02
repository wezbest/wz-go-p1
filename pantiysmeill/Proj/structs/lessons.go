package main

import (
	"fmt"
	"time"
)

// All functions will be using this
type person struct {
	firstName   string
	lastName    string
	contactInfo // Here dont need to declare type
}

// note this is getting embedded in the person struct in the third line -
// This is the meaning of embedding structs
type contactInfo struct {
	email   string
	zipCode int
}

func struc1() {

	// Defining a from the struct person

	// Method 1 - Not good relying on order of the fields
	// Order is interpreted based on the type , not good pracice , because if thiings change then you dont
	// know whoat changed
	momo := person{
		"Mohamed", "Ahmed", contactInfo{"momo@me.com", 99999},
	}

	// Method 2 - Recommended method
	momo2 := person{
		firstName: "Mohamed",
		lastName:  "Ahmed",
	}

	// Method 3  -

	// The zero value is empty sring - {}
	var momo3 person
	momo3.firstName = "Mohamed"
	momo3.lastName = "Ahmed"

	// Printing the stuct variables
	colo("Regarding declaring stucts deifferent methods")
	fmt.Println("\n\n", momo, "\n\n", momo2, "\n\n", momo3, "\n\n")

	// Printing the struct variables - This is better becasse it also prints the labels of he struct
	fmt.Printf("%+v\n", momo3)
}

// Embedding structs

func embstru() {

	tn := time.Now().UTC()

	lina := person{
		firstName: "Lina",
		lastName:  "booty",
		contactInfo: contactInfo{
			email:   "lina@sniffme.com",
			zipCode: 6969,
		},
	}

	colo("Embedding Structs")
	fmt.Println("\n\n", tn)
	fmt.Println("\n\n", lina)
	fmt.Printf("%+v\n", lina)

	lina.updateName("ifa")
	lina.greet()
}

// Update Name

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

// Here we are using the struct above  and then using it with a receiver function
func (p person) greet() {
	colo("Updating stucts")
	fmt.Printf("%+v\n", p)
}
