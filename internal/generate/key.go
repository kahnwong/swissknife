package generate

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/kahnwong/swissknife/internal/utils"
)

func generateKey(n int) (string, error) {
	// https://stackoverflow.com/questions/32349807/how-can-i-generate-a-random-int-using-the-crypto-rand-package/32351471#32351471
	// generate random bytes
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	// convert to base64
	return base64.URLEncoding.EncodeToString(b), nil
}

func Key() error {
	key, err := generateKey(48)
	if err != nil {
		return err
	}

	if err = utils.WriteToClipboard(key); err != nil {
		return err
	}

	fmt.Printf("%s\n", key)
	return nil
}
