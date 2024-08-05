package generate

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/kahnwong/swissknife/utils"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func generateKey(n int) string {
	// https://stackoverflow.com/questions/32349807/how-can-i-generate-a-random-int-using-the-crypto-rand-package/32351471#32351471
	// generate random bytes
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to generate random bytes")
	}

	// convert to base64
	return base64.URLEncoding.EncodeToString(b)
}

var generateKeyCmd = &cobra.Command{
	Use:   "key",
	Short: "Generate key",
	Long:  `Generate key. Re-implementation of "openssl rand -base64 48"". Result is copied to clipboard.`,
	Run: func(cmd *cobra.Command, args []string) {
		key := generateKey(48)
		utils.WriteToClipboard(key)
		fmt.Printf("%s\n", key)
	},
}

func init() {
	Cmd.AddCommand(generateKeyCmd)
}
