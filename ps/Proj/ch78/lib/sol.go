/*
Actual code
*/

package lib

import (
	"fmt"
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
		S1CheckLink(link)
	}

}

func S1CheckLink(link string) {
	response, err := http.Get(link)

	if err != nil {
		fmt.Println(err)
		ChBad("❌", link, "")
		return
	}
	rez := strconv.Itoa(int(response.StatusCode))
	ChGud("✅", link, rez)
}
