/*
Actual code
*/

package lib

import (
	"net/http"
	"strconv"
)

func S1() {

	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://bast.bast",
		"https://bassaat.bast",
	}

	for _, link := range links {
		// Delay for 1 second since this is running sequentially
		S1CheckLink(link)
	}

}

func S1CheckLink(link string) {
	response, err := http.Get(link)

	if err != nil {
		errStr := err.Error() // this is just converting the err project to string
		ChBad("❌", link, errStr)
		return
	}
	rez := strconv.Itoa(int(response.StatusCode))
	ChGud("✅", link, rez)
}
