package post

import "regexp"

func PostalCodeValidator(postalCode string) bool {
	if postalCode == "" {
		return false
	}
	regex := regexp.MustCompile(`[13-9]{4}[1346-9][013-9]{5}`)
	return regex.MatchString(postalCode)
}
