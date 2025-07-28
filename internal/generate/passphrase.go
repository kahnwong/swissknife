package generate

import (
	"fmt"
	"strings"

	"github.com/kahnwong/swissknife/internal/utils"
	"github.com/rs/zerolog/log"
	"github.com/sethvargo/go-diceware/diceware"
)

func generatePassphrase() string {
	// Generate 6 words using the diceware algorithm.
	list, err := diceware.Generate(6)
	if err != nil {
		log.Fatal().Msg("Failed to generate passphrases")
	}

	res := strings.Join(list, "-")

	return res
}

func Passphrase() {
	passphrase := generatePassphrase()
	utils.WriteToClipboard(passphrase)
	fmt.Printf("%s\n", passphrase)
}
