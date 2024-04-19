package iranian_tools

import "github.com/yasinsaee/iranian_tools/util"

// دریافت نام بانک از طریق شماره کارت
func GetBankNameFromCardNumber(cardNumber string) string {
	bankName := util.GetBankNameFromCardNumber(cardNumber)
	return bankName
}
