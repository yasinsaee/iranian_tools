package iranian_tools

import "github.com/YasinSaee/iranian_tools/util"

// تغییر اعداد فارسی به انگلیسی
func ChangeToEnglishDigit(number string) (string, error) {
	return util.ChangeDigit(number, util.English)
}

// تغییر اعداد انگلیسی به فارسی
func ChangeToPersianDigit(number string) (string, error) {
	return util.ChangeDigit(number, util.Persian)
}

// تبدیل اعداد به حروف
func ConvertNumberToWord(number string) (string, error) {
	return util.ConvertNumberToText(number)
}
