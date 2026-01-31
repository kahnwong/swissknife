package generate

import (
	"fmt"

	"github.com/kahnwong/swissknife/internal/utils"
	"github.com/sethvargo/go-password/password"
)

func generatePassword() (string, error) {
	// Generate a password that is 64 characters long with 10 digits, 10 symbols,
	// allowing upper and lower case letters, disallowing repeat characters.
	res, err := password.Generate(32, 10, 0, false, false)
	if err != nil {
		return "", fmt.Errorf("failed to generate password: %w", err)
	}

	return res, nil
}

func Password() error {
	psswd, err := generatePassword()
	if err != nil {
		return err
	}

	if err = utils.WriteToClipboard(psswd); err != nil {
		return err
	}

	fmt.Printf("%s\n", psswd)
	return nil
}
