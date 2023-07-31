/*
Improived sol.go , and to make it concuren. Running concurrency and using
- Channels and go routines
*/

package lib

import (
	"fmt"
	"net/http"
	"time"
)

func S2() {

	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://bast.bast",
		"https://bassaat.bast",
	}

	c := make(chan string) // Make channel

	for _, link := range links {
		go S2CheckLink(link, c)
	}

	// This for loop is for starting a go routine for the S2CheckLink

	for l := range c {
		// go S2CheckLink(l, c)
		go func(link string) { // Anonymous function or funtion literal - fuctions that dont have a name and are called instantly as in the case here
			time.Sleep(3 * time.Second)
			S2CheckLink(link, c)
		}(l)
	}

}

func S2CheckLink(link string, c chan string) {
	response, err := http.Get(link)

	if err != nil {
		// fmt.Println(err)
		errStr := err.Error() // this is just converting the err project to string
		ChBad("❌", link, errStr)
		c <- link
		return
	}

	rez := fmt.Sprintf("%d", response.StatusCode)
	ChGud("✅", link, rez)
	c <- link
}
