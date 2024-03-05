package generate

import (
	"strings"
	"testing"
)

func TestGeneratePassphrase(t *testing.T) {
	passphrase, _ := generatePassphrase()
	if len(strings.Split(passphrase, "-")) != 6 {
		t.Errorf("Result `%s` is not in expected format.", passphrase)
	}
}
