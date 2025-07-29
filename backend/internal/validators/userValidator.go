package validators

import (
	"errors"
	"regexp"
)

func ValidateEmail(email string) error {
	if email == "" {
		return errors.New("email is required")
	}

	var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil
}
