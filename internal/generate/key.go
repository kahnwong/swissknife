package generate

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/kahnwong/swissknife/internal/utils"
	"github.com/rs/zerolog/log"
)

func generateKey(n int) string {
	// https://stackoverflow.com/questions/32349807/how-can-i-generate-a-random-int-using-the-crypto-rand-package/32351471#32351471
	// generate random bytes
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		log.Fatal().Msg("Failed to generate random bytes")
	}

	// convert to base64
	return base64.URLEncoding.EncodeToString(b)
}

func Key() {
	key := generateKey(48)
	utils.WriteToClipboard(key)
	fmt.Printf("%s\n", key)
}
