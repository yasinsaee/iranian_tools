package main

import (
	"time"

	"github.com/YasinSaee/iranian_tools/util/calendar"
)

func main() {

	ti := time.Now()
	t := calendar.FormatIranian(ti)

	// t.Print()
	t.PrintText()
}
