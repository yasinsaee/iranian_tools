package main

import (
	"fmt"

	"github.com/YasinSaee/iranian_tools"
)

func main() {

	// ti := time.Now()
	t, err := iranian_tools.ChangeToEnglishDigit("jhg")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)

	// t.Print()
	// t.PrintText()
}
