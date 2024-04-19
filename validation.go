package iranian_tools

import (
	"github.com/yasinsaee/iranian_tools/util/validators/national"
	"github.com/yasinsaee/iranian_tools/util/validators/phone"
	"github.com/yasinsaee/iranian_tools/util/validators/post"
)

// چک کردن شماره تلفن های ایرانی
func CheckCellPhone(cellPhone string) bool {
	return phone.CellphoneValidator(cellPhone)
}

// چک کردن کد پستی های ایرانی
func CheckPostalCode(postalCode string) bool {
	return post.PostalCodeValidator(postalCode)
}

// چک کردن کد ملی های ایرانی
func CheckNationalCode(nationalCode string) bool {
	return national.NationalCodeValidator(nationalCode)
}
