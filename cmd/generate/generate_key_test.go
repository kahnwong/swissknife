package generate

import (
	"testing"
)

func Test_generateKey(t *testing.T) {
	key, _ := generateKey(48)
	if len(key) != 64 {
		t.Errorf("Result `%s` is not in expected format.", key)
	}
}
