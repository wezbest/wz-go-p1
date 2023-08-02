package files

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

// This is so much smaller than the solution
func Solu() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("bastard: 😡", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)

}

func Solu2() {

	cmd := exec.Command("curl", "https://snips.sh/f/4Gf2ewaj7m")
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		fmt.Println("bastard: 😡", err)
		os.Exit(1)
	}
}

func Solu3() {
	cmd1 := exec.Command("curl", "https://snips.sh/f/4Gf2ewaj7m")

	// Create the second command
	cmd2 := exec.Command("lolcrab", "-g", "inferno")

	// Create a pipe to connect the output of cmd1 to the input of cmd2
	pipe, err := cmd1.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd2.Stdin = pipe

	// Set the output of cmd2 to os.Stdout
	cmd2.Stdout = os.Stdout

	// Start cmd2 before cmd1
	err = cmd2.Start()
	if err != nil {
		log.Fatal(err)
	}

	// Run cmd1
	err = cmd1.Run()
	if err != nil {
		log.Fatal(err)
	}

	// Wait for cmd2 to finish
	err = cmd2.Wait()
	if err != nil {
		log.Fatal(err)
	}

}
