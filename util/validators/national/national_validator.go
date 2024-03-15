package national

import "strconv"

func NationalCodeValidator(nationalCode string) bool {
	if nationalCode == "" {
		return false
	}

	_, err := strconv.Atoi(nationalCode)
	if err != nil {
		return false
	}

	if len(nationalCode) != 10 {
		return false
	}

	nationalCodeLength := 10
	sum := 0

	for i := 0; i < len(nationalCode)-1; i++ {
		digit, err := strconv.Atoi(string(nationalCode[i]))
		if err != nil {
			return false
		}
		sum += digit * nationalCodeLength
		nationalCodeLength--
	}

	remainder := sum % 11
	controlNumber, err := strconv.Atoi(string(nationalCode[9]))
	if err != nil {
		return false
	}

	return (remainder < 2 && controlNumber == remainder) ||
		(remainder >= 2 && controlNumber == 11-remainder)

}
