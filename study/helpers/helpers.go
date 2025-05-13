package helpers

import "strings"

func ValidateUserName(userFirstName string) bool {
	return len(userFirstName) >= 2
}

func ValidateUserLastName(userLastName string) bool {
	return len(userLastName) >= 2
}

func ValidateUserEmail(userEmail string) bool {
	return strings.Contains(userEmail, "@")
}
