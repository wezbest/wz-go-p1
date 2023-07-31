// This is fetching image from
//https://snips.sh/f/wtwvZStcRO

package lib

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func Bal() {
	cmd := exec.Command("curl", "https://snips.sh/f/wtwvZStcRO")
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println("bastard: 😡", err)
		os.Exit(1)
	}
}

func BaWrite() {

	cmd := exec.Command("curl", "https://snips.sh/f/wtwvZStcRO")

	// Create file
	f, err := os.Create("output.fish")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Set command stdout to both file and os.Stdout
	cmd.Stdout = io.MultiWriter(f, os.Stdout)

	err = cmd.Run()
	if err != nil {
		panic(err)
	}

}
