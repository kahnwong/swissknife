package generate

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log/slog"

	"github.com/kahnwong/swissknife/cmd/utils"
	"github.com/spf13/cobra"
)

// https://stackoverflow.com/questions/32349807/how-can-i-generate-a-random-int-using-the-crypto-rand-package/32351471#32351471
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateKey(s int) (string, error) {
	b, err := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

var generateKeyCmd = &cobra.Command{
	Use:   "key",
	Short: "Generate key",
	Long:  `Generate key. Re-implementation of "openssl rand -base64 48"". Result is copied to clipboard.`,
	Run: func(cmd *cobra.Command, args []string) {
		// main
		key, err := generateKey(48)
		if err != nil {
			slog.Error("Error generating key")
		}
		utils.WriteToClipboard(key)
		fmt.Printf("%s\n", key)
	},
}

func init() {
	Cmd.AddCommand(generateKeyCmd)
}
