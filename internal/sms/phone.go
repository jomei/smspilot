package sms

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func PhoneGet(phoneStr string) (int, error) {
	p := clearStringPhone(phoneStr)
	if !isPhone(p) {
		return 0, errors.New("bad number phone")
	}
	phoneInt, err := ConvertPhoneToInt(p)
	if err != nil {
		return 0, err
	}
	return phoneInt, nil
}

func ConvertPhoneToInt(phone string) (int, error) {
	phoneInt, err := strconv.Atoi(phone)
	if err != nil {
		fmt.Println("error: strconv.Atoi(phone)")
		return 0, err
	}
	return phoneInt, nil
}

func isPhone(phone string) bool {

	if len(phone) != 12 {
		return false
	}
	if !strings.HasPrefix(phone, "998") {
		return false
	}
	return true
}

func clearStringPhone(str string) string {
	var r = regexp.MustCompile(`[^0-9]+`)
	return r.ReplaceAllString(str, "")
}
