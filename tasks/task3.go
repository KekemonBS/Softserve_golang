package tasks

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Task3(creditCardNumber string) (isValid bool, lastFourDigits string) {
	isValid, err := regexp.Match(`^([0-9]{4}(\s*|)){4}`, []byte(creditCardNumber))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return isValid, strings.ReplaceAll(creditCardNumber, " ", "")[12:]
}
