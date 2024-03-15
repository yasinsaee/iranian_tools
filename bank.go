package iranian_tools

import "github.com/YasinSaee/iranian_tools/util"

func GetBankNameFromCardNumber(cardNumber string) string {
	bankName := util.GetBankNameFromCardNumber(cardNumber)
	return bankName
}
