package format

import "strings"

// PhoneFormatCommon ...
func PhoneFormatCommon(phone string) string {
	// Check phone is null
	if phone == "" {
		return "Vui lòng nhập số điện thoại vào!"
	}

	// Split all space
	phone = strings.Replace(phone, " ", "", -1)

	// If first character is "0", replace to "+84", become "+84..."
	if string(phone[0]) == "0" {
		phone = strings.Replace(phone, "0", "+84", 1)
	}

	// If 2 first characters is "84", add "+"
	if phone[0:2] == "84" {
		phone = "+" + phone
	}

	return phone
}
