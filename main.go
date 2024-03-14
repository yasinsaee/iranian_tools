package main

import (
	"iranian_tools/util/calendar"
	"time"
)

func main() {

	ti := time.Now()
	t := calendar.FormatIranian(ti)

	// t.Print()
	t.PrintText()
}
