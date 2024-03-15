package main

import (
	"fmt"
	"time"

	"github.com/YasinSaee/iranian_tools"
)

func main() {

	t, err := iranian_tools.ChangeToEnglishDigit("۱۲۳۴۵۶")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)

	//----------------------------------------------------------------

	ok := iranian_tools.CheckCellPhone("09039686577")
	fmt.Println(ok)

	//----------------------------------------------------------------

	ti := time.Now()
	time := iranian_tools.ChangeToJalali(ti)
	fmt.Println(time.Year)

}
