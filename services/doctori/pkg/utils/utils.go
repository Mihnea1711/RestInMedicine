package utils

import "regexp"

var (
	EmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	PhoneRegex = regexp.MustCompile(`^(07[0-9]{8}|\+407[0-9]{8})$`)
)
