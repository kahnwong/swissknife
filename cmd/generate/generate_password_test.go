package generate

import (
	"regexp"
	"strings"
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	password, _ := generatePassword()

	// length must be 32
	if len(password) != 32 {
		t.Errorf("Result `%s` does not have required length of 32.", password)
	}

	// must contain 10 digits
	re := regexp.MustCompile("[0-9]+")
	numberMatches := re.FindAllString(password, -1)
	if len(strings.Join(numberMatches, "")) != 10 {
		t.Errorf("Result `%s` does not contain 10 digits.", password)
	}
}
