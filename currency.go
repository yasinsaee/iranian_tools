package iranian_tools

import "github.com/YasinSaee/iranian_tools/util"

// تبدیل قیمت از تومان به ریال
// یا برعکس از ریال به تومان
func ConvertCurrency(oldCurrency, newCurrency string, amount float64) (util.CustomFloat64, error) {
	return util.ConvertCurrency(oldCurrency, newCurrency, amount)
}
