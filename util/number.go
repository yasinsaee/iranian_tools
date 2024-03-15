package util

import "strings"

type NumStrLanguage int

const (
	English NumStrLanguage = iota
	Persian
)

func ChangeToEnglishDigit(number string) string {
	return changeDigit(number, English)
}

func ChangeToPersianDigit(number string) string {
	return changeDigit(number, Persian)
}

func changeDigit(input string, lang NumStrLanguage) string {
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
