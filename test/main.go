package main

import (
	"fmt"

	"github.com/YasinSaee/iranian_tools"
)

func main() {

	// englishNumber, _ := iranian_tools.ChangeToEnglishDigit("۱۲۳۴۵۶")
	// fmt.Println(englishNumber) // Output: 123456

	//----------------------------------------------------------------

	// ok := iranian_tools.CheckCellPhone("09030000000")
	// fmt.Println(ok) // Output: true or false
	//----------------------------------------------------------------

	// ti := time.Now()
	// jalaliDate := iranian_tools.ChangeToJalali(ti)
	// fmt.Printf("%d/%02d/%02d", jalaliDate.Year, jalaliDate.Month, jalaliDate.Day) // Output: 1402/12/25

	am, _ := iranian_tools.ConvertCurrency(iranian_tools.TomanCurrency, iranian_tools.RialCurrency, 560000)
	fmt.Println(am.SeparateDigits(4))
}
