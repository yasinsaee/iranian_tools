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

To use IranianTools in your Go project, you can simply import it:

```bash
go get github.com/YasinSaee/iranian_tools
```

# Usage Examples
  - ## ConvertToEnglishDigit
    This function converts a string containing Persian digits to a string containing English digits.
    Example:
    ```go
    englishNumber, _ := iranian_tools.ChangeToEnglishDigit("۱۲۳۴۵۶")
    ```
  - ## ConvertToPersianDigit
    This function converts a string containing English digits to a string containing Persian digits.
    Example:
    ```go
    persianNumber, _ := iranian_tools.ChangeToPersianDigit("123456")
    ```
  - ## ConvertNumberToPersianLetters
    This function converts a string containing digits to their equivalent words in Persian.
    Example:
    ```go
    persianWord, _ := iranian_tools.ConvertNumberToWord("154")
    ```



