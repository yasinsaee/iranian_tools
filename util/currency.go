//@author:YasinSaee

package util

import "fmt"

// تبدیل تومان به ریال,
// تبدیل ریال به تومان
func ConvertCurrency(oldCurrency, newCurrency string, amount float64) (CustomFloat64, error) {
	conversionRates := map[string]float64{
		"rial_to_toman": 0.1,
		"toman_to_rial": 10.0,
	}

	rateName := fmt.Sprintf("%s_to_%s", oldCurrency, newCurrency)
	conversionRate, found := conversionRates[rateName]
	if !found {
		return 0, fmt.Errorf("can not convert")
	}

	convertedAmount := amount / conversionRate

	return CustomFloat64(convertedAmount), nil
}
