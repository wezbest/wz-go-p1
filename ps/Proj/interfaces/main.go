/*

Interfaces Lessons - This will be out interfaces which are important aspect of thelang
Actual code for both of these are in les.go , Here we are just callin the function directly , but all variable
and funtional definitions are in les.go

*/

package main

func main() {
	colo("Interfaces")

	hea("English Bot Print")
	eb := englishBot{}
	printGreeting(eb)

	hea("Bootish Bot Print")
	bb := bootishBot{}
	printGreeting(bb)
	
}
