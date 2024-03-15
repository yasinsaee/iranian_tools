package iranian_tools

import "github.com/YasinSaee/iranian_tools/util"

func ChangeToEnglishDigit(number string) string {
	return util.ChangeDigit(number, util.English)
}

func ChangeToPersianDigit(number string) string {
	return util.ChangeDigit(number, util.Persian)
}
