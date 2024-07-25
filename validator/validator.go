package validator

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUserName = regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\\s]+$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("length must be between %d and %d", minLength, maxLength)
	}
	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 25); err != nil {
		return err
	}

	if !isValidUserName(value) {
		return fmt.Errorf("username can only contain letters, numbers and underscores")
	}

	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 8, 50)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 5, 50); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("invalid email address")
	}

	return nil
}

func ValidateFullName(value string) error {
	if err := ValidateString(value, 3, 25); err != nil {
		return err
	}

	if !isValidFullName(value) {
		return fmt.Errorf("must contain only letters and spaces")
	}

	return nil
}
