# IranianTools

IranianTools is a versatile Go package designed for developers working with the Go programming language. It provides various functionalities to facilitate common tasks encountered in Iranian software development.

## Features

- **Number to Words Conversion**: Convert numeric values to their equivalent words in Persian or English.
- **Phone Number Validation**: Validate Iranian phone numbers.
- **Postal Code Validation**: Validate Iranian postal codes.
- **National ID Validation**: Validate Iranian national identification numbers (Melli Code).
- **Bank Name Recognition**: Recognize the bank name from a given credit card number.
- **Convert Gregorian Year to Jalali (Shamsi) Year**: Convert Gregorian calendar years to Jalali years and vice versa.
- ** and more ...!**

## Installation

To use IranianTools in your Go project, you can simply import it:

```bash
go get github.com/YasinSaee/iranian_tools
```

## Usage Examples 
- ### Number
  - ### ConvertToEnglishDigit
    This function converts a string containing Persian digits to a string containing English digits.
    
    Example:
    ```go
    englishNumber, _ := iranian_tools.ChangeToEnglishDigit("۱۲۳۴۵۶")
    fmt.Println(englishNumber) // Output: 123456
    ```
  - ### ConvertToPersianDigit
    This function converts a string containing English digits to a string containing Persian digits.
    
    Example:
    ```go
    persianNumber, _ := iranian_tools.ChangeToPersianDigit("123456")
    fmt.Println(persianNumber) // Output: ۱۲۳۴۵۶
    ```
  - ### ConvertNumberToPersianLetters
    This function converts a string containing digits to their equivalent words in Persian.
    
    Example:
    ```go
    persianWord, _ := iranian_tools.ConvertNumberToWord("154")
    fmt.Println(persianWord) // Output: صد و پنجاه و چهار
    ```

- ### Validations

  - ### CheckCellPhone
    This function validates Iranian phone numbers.
    
    Example:
    ```go
    ok := iranian_tools.CheckCellPhone("09030000000")
    fmt.Println(ok) // Output: true or false
    ```
    
  - ### CheckPostalCode
    This function validates Iranian postal codes.
    
    Example:
    ```go
    ok := iranian_tools.CheckPostalCode("123456789") // Please set valid postalCode
    fmt.Println(ok) // Output: true or false
    ```
    
  - ### CheckNationalCode
    This function validates Iranian national codes.
    
    Example:
    ```go
    ok := iranian_tools.CheckNationalCode("5858960707")
    fmt.Println(ok) // Output: true or false
    ```
- ### Calendar
  - ### ChangeToJalali
     Converts a Gregorian date to Jalali (Iranian) calendar.
      
     Example :
     ```go
     ti := time.Now()
     jalaliDate := iranian_tools.ChangeToJalali(ti)
     fmt.Printf("%d/%02d/%02d", jalaliDate.Year, jalaliDate.Month, jalaliDate.Day) // Output: 1402/12/25
     ```

## Contribution

Contributions are welcome! If you have any suggestions, bug reports, or feature requests, please open an issue or submit a pull request.