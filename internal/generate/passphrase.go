package generate

import (
	"fmt"
	"strings"

	"github.com/kahnwong/swissknife/internal/utils"
	"github.com/sethvargo/go-diceware/diceware"
)

func generatePassphrase() (string, error) {
	// Generate 6 words using the diceware algorithm.
	list, err := diceware.Generate(6)
	if err != nil {
		return "", fmt.Errorf("failed to generate passphrases: %w", err)
	}

	res := strings.Join(list, "-")

	return res, nil
}

func Passphrase() error {
	passphrase, err := generatePassphrase()
	if err != nil {
		return err
	}

	if err = utils.WriteToClipboard(passphrase); err != nil {
		return err
	}

	fmt.Printf("%s\n", passphrase)
	return nil
}
