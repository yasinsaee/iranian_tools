package validators

import "strconv"

type NationalCodeValidator struct {
	NationalCode string `json:"national_code"`
}

func (v *NationalCodeValidator) Validate() bool {
	if v.NationalCode == "" {
		return false
	}

	_, err := strconv.Atoi(v.NationalCode)
	if err != nil {
		return false
	}

	if len(v.NationalCode) != 10 {
		return false
	}

	nationalCodeLength := 10
	sum := 0

	for i := 0; i < len(v.NationalCode)-1; i++ {
		digit, err := strconv.Atoi(string(v.NationalCode[i]))
		if err != nil {
			return false
		}
		sum += digit * nationalCodeLength
		nationalCodeLength--
	}

	remainder := sum % 11
	controlNumber, err := strconv.Atoi(string(v.NationalCode[9]))
	if err != nil {
		return false
	}

	return (remainder < 2 && controlNumber == remainder) ||
		(remainder >= 2 && controlNumber == 11-remainder)

}
