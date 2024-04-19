package iranian_tools

import (
	"strconv"

	"github.com/yasinsaee/iranian_tools/util"
)

// تغییر اعداد فارسی به انگلیسی
func ChangeToEnglishDigit(number string) (string, error) {
	return util.ChangeDigit(number, util.English)
}

// تغییر اعداد انگلیسی به فارسی
func ChangeToPersianDigit(number string) (string, error) {
	return util.ChangeDigit(number, util.Persian)
}

// تبدیل اعداد به حروف فارسی
func ConvertNumberToWord(number string) (string, error) {
	return util.ConvertNumberToText(number)
}

// .جدا کردن اعداد از هم به صورت کاستوم
// divLen عدد شما را به صورت دلخواه می تواند تغییر دهد.
// به عنوان مثال ۳ رقم رقم یا ۲ رقم
func SeparateDigits(number interface{}, divLen int) string {
	switch num := number.(type) {
	case int:
		strNum := strconv.Itoa(num)
		return util.SeparateDigits(strNum, divLen)
	case float64:
		strNum := strconv.FormatFloat(num, 'f', -1, 64)
		return util.SeparateDigits(strNum, divLen)
	case util.CustomFloat64:
		{
			strNum := strconv.FormatFloat(float64(num), 'f', -1, 64)
			return util.SeparateDigits(strNum, divLen)
		}
	case util.CustomInt:
		{
			strNum := strconv.Itoa(int(num))
			return util.SeparateDigits(strNum, divLen)
		}
	default:
		return "Unsupported type"
	}
}
