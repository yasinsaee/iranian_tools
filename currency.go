package iranian_tools

import "github.com/YasinSaee/iranian_tools/util"

const (
	RialCurrency  = "rial"
	TomanCurrency = "toman"
)

func ConvertCurrency(oldCurrency, newCurrency string, amount float64) (util.CustomFloat64, error) {
	return util.ConvertCurrency(oldCurrency, newCurrency, amount)
}
