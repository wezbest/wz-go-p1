/*
Actual lessons code all go in here
- Writing the bots progrm hers as two seperae functions here where we will call it in main.gol
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func ht1() {
	hea("Fist function - Getting HTTP respons from Google ")
	resp, err := http.Get("http://example.com")
	if err != nil {
		fmt.Println("😡:", err)
		os.Exit(1)
		panic(err)
	}
	fmt.Println(resp)
}

func ht2() {
	hea("Second function - Working with red function - USe of the make function fetching th html from page ")
	hea("Ping http://example.com ")
	resp, err := http.Get("http://example.com")
	if err != nil {
		fmt.Println("😡:", err)
		os.Exit(1)
		panic(err)
	}

	// Make function - in The following function it will make a byte sluce with 99999 elements in it
	// Read funtioon is incapable of resizing the byte slice, so we declare the max space for it
	bs := make([]byte, 99999)
	resp.Body.Read(bs)

	fmt.Println(string(bs))
	// fmt.Println(bs) - thi function will dump all zeroes from byte slice which cant be read so conersion to string is making it buman readable

}

/*
Function 3- This is condensing down code on top , we are reducing the lines of code, insread of
fmt.println() , we are using now io.Copy()
*/

func ht3() {
	hea("Third function - using io.Copy(os.Stdout, resp.body) ")
	hea("Ping http://example.com ")
	resp, err := http.Get("http://example.com")
	if err != nil {
		fmt.Println("😡:", err)
		os.Exit(1)
		panic(err)
	}

	// This is using both the reader and writer interface - https://pkg.go.dev/io#Copy
	// This also means at this point you should also be able to write to a file
	// this was added by you and it works
	io.Copy(os.Stdout, resp.Body)
}

// Found this on claude for writing the respons to a file
func ht4() {

	hea("4th function - Use Claude the above output it going to be written to a file")

	resp, err := http.Get("http://example.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Create file to write to
	file, err := os.Create("asslick.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// Write response body to file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Licked her booty")

}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

/*
Writing you won writer interface
*/

// This is required for implementing the writer interface
type logWriter struct {
}

func ht6() {
	hea("6th function - Implementing our own writer interface")
	hea("Ping http://example.com ")
	resp, err := http.Get("http://example.com")
	if err != nil {
		fmt.Println("😡:", err)
		os.Exit(1)
		panic(err)
	}

	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes", len(bs))
	return len(bs), nil
}

//++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
