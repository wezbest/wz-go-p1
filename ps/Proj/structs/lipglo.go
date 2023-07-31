// Attempting to make a lipglo function

package  main

import (
	lg "github.com/charmbracelet/lipgloss"
	f "fmt"
)

/*
The following fnction is super sexy
1. First of all you are making a simple function that has a receiver of type string
2. Then you are doing all the coloring with a lipgloss style
3. Note that since you put a receiver of type string, you are passing in a string
4. Its part of packge main , so you can now call it in any file amazingly
5. Then you just call it in any function ensuring that what is being passed in is a string
6. You can now call this function anywhere
sna
*/

func colo(s string) {
	var liq = lg.NewStyle().
		Foreground(lg.Color("#F6FA70")).
		Background(lg.Color("#B3005E")).
		Width(50).Align(lg.Center).
		Blink(true).Border(lg.DoubleBorder(), true, true, true, true).
		PaddingTop(1).
		PaddingBottom(1).
		BorderForeground(lg.Color("#F6FA70"))

	f.Println(liq.Render(s))

}