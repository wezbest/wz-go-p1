// Attempting to make a lipglo function

package lip

import (
	f "fmt"
	"time"
	lg "github.com/charmbracelet/lipgloss"
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

var liq = lg.NewStyle().
	Foreground(lg.Color("#FF55BB")).
	Background(lg.Color("#060047")).
	Width(70).Align(lg.Center).Bold(true).
	Blink(true).Border(lg.DoubleBorder(), true, true, true, true).
	PaddingTop(1).
	PaddingBottom(1).
	BorderForeground(lg.Color("#F6FA70"))

var banner = `


╱╭━━━╮╱╭━╮╭━╮╱╭━━━╮╱╭━━━╮╱╭━━━╮╱╭━━━╮╱╭━━╮╱╭━━━╮╱╭━━━╮
╱┃╭━━╯╱╰╮╰╯╭╯╱┃╭━╮┃╱┃╭━━╯╱┃╭━╮┃╱┃╭━╮┃╱╰┫┣╯╱┃╭━╮┃╱┃╭━━╯
╱┃╰━━╮╱╱╰╮╭╯╱╱┃┃╱╰╯╱┃╰━━╮╱┃╰━╯┃╱┃┃╱╰╯╱╱┃┃╱╱┃╰━━╮╱┃╰━━╮
╱┃╭━━╯╱╱╭╯╰╮╱╱┃┃╱╭╮╱┃╭━━╯╱┃╭╮╭╯╱┃┃╱╭╮╱╱┃┃╱╱╰━━╮┃╱┃╭━━╯
╱┃╰━━╮╱╭╯╭╮╰╮╱┃╰━╯┃╱┃╰━━╮╱┃┃┃╰╮╱┃╰━╯┃╱╭┫┣╮╱┃╰━╯┃╱┃╰━━╮
╱╰━━━╯╱╰━╯╰━╯╱╰━━━╯╱╰━━━╯╱╰╯╰━╯╱╰━━━╯╱╰━━╯╱╰━━━╯╱╰━━━╯


`

func Colo(s string) {
	t := time.Now().UTC().Format("Mon Jan 02 15:04:05 UTC")

	f.Println(liq.Render(banner, t, "\n\n", s))

}

var lqh = lg.NewStyle().
	Background(lg.Color("#0C1E7F")).
	Foreground(lg.Color("#FF008E")).
	Bold(true).
	Underline(true)

func Hea(s string) {
	f.Println("\n")
	f.Println(lqh.Render(s), "\n")
}

var texco = lg.NewStyle().
	Background(lg.Color("#3A1078")).
	Foreground(lg.Color("#00DFA2"))

func Texc(s string, t string) {
	f.Println(texco.Render(s, t))
}
