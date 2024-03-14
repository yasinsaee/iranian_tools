package util

import (
	"strings"
	"time"
)

func GetIranianMonthLetter(monthNumber string) string {
	monthNumber = ChangeDigit(monthNumber, English)
	outputText := ""
	if string(monthNumber[0]) == "0" {
		monthNumber = strings.ReplaceAll(monthNumber, "0", "")
	}
	switch monthNumber {
	case "1":
		outputText = "فروردین"
	case "2":
		outputText = "اردیبهشت"
	case "3":
		outputText = "خرداد"
	case "4":
		outputText = "تیر"
	case "5":
		outputText = "مرداد"
	case "6":
		outputText = "شهریور"
	case "7":
		outputText = "مهر"
	case "8":
		outputText = "ابان"
	case "9":
		outputText = "اذر"
	case "10":
		outputText = "دی"
	case "11":
		outputText = "بهمن"
	case "12":
		outputText = "اسفند"
	}
	return outputText
}

type NumStrLanguage int

const (
	English NumStrLanguage = iota
	Persian
)

func ChangeDigit(input string, lang NumStrLanguage) string {
	if lang == English {
		input = strings.ReplaceAll(input, "۰", "0")
		input = strings.ReplaceAll(input, "۱", "1")
		input = strings.ReplaceAll(input, "۲", "2")
		input = strings.ReplaceAll(input, "۳", "3")
		input = strings.ReplaceAll(input, "۴", "4")
		input = strings.ReplaceAll(input, "۵", "5")
		input = strings.ReplaceAll(input, "۶", "6")
		input = strings.ReplaceAll(input, "۷", "7")
		input = strings.ReplaceAll(input, "۸", "8")
		input = strings.ReplaceAll(input, "۹", "9")
	} else if lang == Persian {
		input = strings.ReplaceAll(input, "0", "۰")
		input = strings.ReplaceAll(input, "1", "۱")
		input = strings.ReplaceAll(input, "2", "۲")
		input = strings.ReplaceAll(input, "3", "۳")
		input = strings.ReplaceAll(input, "4", "۴")
		input = strings.ReplaceAll(input, "5", "۵")
		input = strings.ReplaceAll(input, "6", "۶")
		input = strings.ReplaceAll(input, "7", "۷")
		input = strings.ReplaceAll(input, "8", "۸")
		input = strings.ReplaceAll(input, "9", "۹")
	}
	return input
}

// نمایش اختلاف بین دو تاریخ به صورت خلاصه وار و متنی
func DifferenceDate(inputDatetime time.Time) string {
	systemDatetime := time.Now()

	differenceInDays := int(systemDatetime.Sub(inputDatetime).Hours() / 24)
	if differenceInDays > 7 {
		if differenceInDays > 7 && differenceInDays < 14 {
			return "هفته پیش"
		} else if differenceInDays < 7*4 {
			return "چند هفته پیش"
		} else if differenceInDays >= 30 && differenceInDays < 60 {
			return "ماه پیش"
		} else if differenceInDays < 365 {
			return "چند ماه پیش"
		} else if differenceInDays >= 365 && differenceInDays < 730 {
			return "سال پیش"
		} else {
			return "چند سال پیش"
		}
	} else {
		differenceInSeconds := int(systemDatetime.Sub(inputDatetime).Seconds())
		if differenceInSeconds < 60 {
			return "لحظاتی پیش"
		} else if differenceInSeconds < 60*60 {
			return "دقایقی پیش"
		} else if differenceInSeconds < 3600*24 {
			return "ساعاتی پیش"
		} else if differenceInDays == 1 {
			return "دیروز"
		} else {
			return "چند روز پیش"
		}
	}
}
