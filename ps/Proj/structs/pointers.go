// Pointers learning

package main

import "fmt"

type per1 struct {
	firstName string
	lastName  string
	per2
}

type per2 struct {
	email   string
	zipCode int
}

func poi1() {

	moina := per1{
		firstName: "Lina",
		lastName:  "booty",
		per2: per2{
			email:   "moina@sniffme.com",
			zipCode: 6969,
		},
	}

	colo("Pointer Study Function")
	fmt.Printf("%+v\n", moina)

	// moinaPointer := &moina // convert to memory address
	// fmt.Println(moinaPointer)
	// moinaPointer.updateName("ifa")
	moina.updateName("ifa")
	moina.print()
}

// Complex pointer
func (pointerToPer1 *per1) updateName(newFirstName string) {
	(*pointerToPer1).firstName = newFirstName // change it to a value
}

func (p per1) print() {
	fmt.Printf("%+v\n", p)
}

// Pass by value function - Copy all data in struct and then make new one in memory

// func (p per1) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }
