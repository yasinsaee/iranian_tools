# IranianTools

IranianTools is a versatile Go package designed for developers working with the Go programming language. It provides various functionalities to facilitate common tasks encountered in Iranian software development.

## Features

- **Number to Words Conversion**: Convert numeric values to their equivalent words in Persian or English.
- **Phone Number Validation**: Validate Iranian phone numbers.
- **Postal Code Validation**: Validate Iranian postal codes.
- **National ID Validation**: Validate Iranian national identification numbers (Melli Code).
- **Bank Name Recognition**: Recognize the bank name from a given credit card number.
- **Convert Gregorian Year to Jalali (Shamsi) Year**: Convert Gregorian calendar years to Jalali years and vice versa.
- **... and more!**

# Installation

go get github.com/YasinSaee/iranian_tools

## Usage Examples
### ConvertToEnglishDigit

This function converts a string containing Persian digits to a string containing English digits.

Example:
```go
t, err := iranian_tools.ChangeToEnglishDigit("۱۲۳۴۵۶")
if err != nil {
    fmt.Println(err)
}
fmt.Println(t)

# Contribution

Contributions are welcome! If you have any suggestions, bug reports, or feature requests, please open an issue or submit a pull request.
