/*
This is the hint which was provided regarding getting arguments
*/

package files

import (
	"fmt"
	"har/src"
	"os"
)

func Hint1() {
	fmt.Println(os.Args)
	src.Para("Printing os.Args[0]")
	fmt.Println(os.Args[0])
	src.Para("Printing os.Args[1]")
	fmt.Println(os.Args[1])
}
