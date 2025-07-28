package generate

import (
	"fmt"

	"github.com/kahnwong/swissknife/internal/utils"
	"github.com/rs/zerolog/log"
	"github.com/sethvargo/go-password/password"
)

func generatePassword() string {
	// Generate a password that is 64 characters long with 10 digits, 10 symbols,
	// allowing upper and lower case letters, disallowing repeat characters.
	res, err := password.Generate(32, 10, 0, false, false)
	if err != nil {
		log.Fatal().Msg("Failed to generate password")
	}

	return res
}

func Password() {
	psswd := generatePassword()
	utils.WriteToClipboard(psswd)
	fmt.Printf("%s\n", psswd)
}
