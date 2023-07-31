/*
Thisis the code directly taken from
https://github.com/charmbracelet/lg/tree/master/examples/layout
*/

package src

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	lg "github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var (
	dialogBoxStyle = lg.NewStyle().
			Border(lg.RoundedBorder()).
			BorderForeground(lg.Color("#874BFD")).
			Padding(1, 0).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true)

	buttonStyle = lg.NewStyle().
			Foreground(lg.Color("#FFF7DB")).
			Background(lg.Color("#888B7E")).
			Padding(0, 3).
			MarginTop(1)

	activeButtonStyle = buttonStyle.Copy().
				Foreground(lg.Color("#FFF7DB")).
				Background(lg.Color("#F25D94")).
				MarginRight(2).
				Underline(true)

	width = 96

	docStyle = lipgloss.NewStyle().Padding(1, 2, 1, 2)
)

func Di() {
	// Dialog

	physicalWidth, _, _ := term.GetSize(int(os.Stdout.Fd()))
	doc := strings.Builder{}

	okButton := activeButtonStyle.Render("Yes")
	cancelButton := buttonStyle.Render("Maybe")

	question := lg.NewStyle().Width(50).Align(lg.Center).Render("Are you sure you want to eat marmalade?")
	buttons := lg.JoinHorizontal(lg.Top, okButton, cancelButton)
	ui := lg.JoinVertical(lg.Center, question, buttons)

	dialog := lg.Place(width, 9,
		lg.Center, lg.Center,
		dialogBoxStyle.Render(ui),
		lg.WithWhitespaceChars("猫咪"),
	)

	if physicalWidth > 0 {
		docStyle = docStyle.MaxWidth(physicalWidth)
	}

	doc.WriteString(dialog + "\n\n")
	fmt.Println(docStyle.Render(doc.String()))

}
