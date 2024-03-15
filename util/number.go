//@author:YasinSaee

package util

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type (
	NumStrLanguage int
	CustomFloat64  float64
	CustomInt      int
	CustomString   string
)

const (
	English NumStrLanguage = iota
	Persian
)

// این فانکشن می تواند اعداد انگلیسی را به فارسی
// و برعکس آن یعنی از فارسی به انگلیسی تبدیل کند.
func ChangeDigit(input string, lang NumStrLanguage) (string, error) {
	if !IsNumeric(input) {
		return "", errors.New("please use number")
	}
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
	return input, nil
}

// چک کردن متن که حتما به صورت اعداد فارسی یا انگلیسی وارد شده باشد.
func IsNumeric(s string) bool {
	match, _ := regexp.MatchString("^[0-9۰-۹]*$", s)
	return match
}

var units = []string{
	"",
	"یک",
	"دو",
	"سه",
	"چهار",
	"پنج",
	"شش",
	"هفت",
	"هشت",
	"نه",
	"ده",
	"یازده",
	"دوازده",
	"سیزده",
	"چهارده",
	"پانزده",
	"شانزده",
	"هفده",
	"هجده",
	"نوزده",
}

var tens = []string{
	"",
	"ده",
	"بیست",
	"سی",
	"چهل",
	"پنجاه",
	"شصت",
	"هفتاد",
	"هشتاد",
	"نود",
}

var hundreds = []string{
	"",
	"صد",
	"دویست",
	"سیصد",
	"چهارصد",
	"پانصد",
	"ششصد",
	"هفتصد",
	"هشتصد",
	"نهصد",
}

var chunks = []string{
	"",
	"هزار",
	"میلیون",
	"میلیارد",
	"بیلیون",
	"بیلیارد",
	"تریلیون",
	"تریلیارد",
	"کوآدریلیون",
	"کادریلیارد",
	"کوینتیلیون",
	"کوانتینیارد",
	"سکستیلیون",
	"سکستیلیارد",
	"سپتیلیون",
	"سپتیلیارد",
	"اکتیلیون",
	"اکتیلیارد",
	"نانیلیون",
	"نانیلیارد",
	"دسیلیون",
}

func threeDigitToPersianWord(number int) string {
	result := make([]string, 0, 3)
	chunkUnits := number % 10
	chunkTens := (number / 10) % 10
	chunkHundreds := (number / 100) % 10

	if chunkHundreds != 0 {
		result = append(result, hundreds[chunkHundreds])
	}

	if chunkTens == 1 {
		result = append(result, units[chunkUnits+10])
	} else {
		if chunkTens != 0 {
			result = append(result, tens[chunkTens])
		}
		if chunkUnits != 0 {
			result = append(result, units[chunkUnits])
		}
	}

	return strings.Join(result, " و ")
}

// تبدیل اعداد به حروف فارسی
func ConvertNumberToText(number string) (string, error) {
	nm, err := ChangeDigit(number, English)
	if err != nil {
		return "", err
	}

	num, _ := strconv.Atoi(nm)

	if num == 0 {
		return "صفر", nil
	}

	negative := false

	if num < 0 {
		negative = true
		num *= -1
	}

	count := 0
	result := make([]string, 0, 3)

	for num > 0 {
		chunk := num % 1000
		if chunk == 0 {
			num /= 1000
			count++
			continue
		}

		chunkWords := threeDigitToPersianWord(chunk)
		if chunk != 0 {
			result = append([]string{chunkWords + " " + chunks[count]}, result...)
		} else {
			result = append([]string{chunkWords}, result...)
		}
		num /= 1000
		count++
	}

	resultStr := strings.Join(result, " و ")
	resultStr = strings.Trim(resultStr, " ")

	if negative {
		resultStr = "منفی " + resultStr
	}

	return resultStr, nil
}

func (number *CustomInt) SeparateDigits(divLen int) string {
	strNum := strconv.Itoa(int(*number))
	return SeparateDigits(strNum, divLen)
}

func (number *CustomFloat64) SeparateDigits(divLen int) string {
	strNum := strconv.FormatFloat(float64(*number), 'f', -1, 64)
	return SeparateDigits(strNum, divLen)
}

func SeparateDigits(strNum string, divLen int) string {
	decimalIndex := -1
	for i, char := range strNum {
		if char == '.' {
			decimalIndex = i
			break
		}
	}

	if decimalIndex == -1 {
		decimalIndex = len(strNum)
	}

	result := ""

	digitCounter := 0

	for i := decimalIndex - 1; i >= 0; i-- {
		result = string(strNum[i]) + result

		digitCounter++

		if digitCounter == divLen && i != 0 {
			result = "," + result
			digitCounter = 0
		}
	}

	return result
}
