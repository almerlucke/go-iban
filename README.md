### Basic Golang IBAN validator

Validations:
- Validate if country code and check digits are present
- Validate if country code is in accepted country code list
- Validate national BBAN code format
- Validate mod97 check digits

#### Testing
Test suite includes example of IBAN's for most countries and fake IBAN's. Errors are as specific as possible.

#### Example

	package main

	import (
		"fmt"
		"github.com/almerlucke/go-iban/iban"
	)
	
	func main() {
		iban, err := iban.NewIBAN("NL40ABNA0517552264")
	
		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			fmt.Printf("%v\n", iban.PrintCode)
			fmt.Printf("%v\n", iban.Code)
		}
	}