package main

// Testng lipgloss finally here insted os using a anoher color repo
// This one has many more cool functions as compared to the rest

import (
	f "fmt"
	"time"

	lg "github.com/charmbracelet/lipgloss"
)

var style = lg.NewStyle().
	Bold(true).
	Foreground(lg.Color("#FAFAFA")).
	Background(lg.Color("#151B54")).
	PaddingTop(2).
	PaddingBottom(2).
	Width(22).
	Align(lg.Center).
	Blink(true).
	Border(lg.DoubleBorder(), true, true, true, true)

func colop() {
	t := time.Now()

	f.Println(style.Render("Lick it Good \n\n", t.String())) // Here we converting the time to string
	f.Println(style.Render("t"))
}
