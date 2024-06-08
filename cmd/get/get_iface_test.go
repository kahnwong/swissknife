package get

import (
	"strings"
	"testing"
)

func TestGetIface(t *testing.T) {
	iface, err := getIface()
	if err != nil {
		t.Errorf("getIface() error = %v", err)
	} else {
		isEn := strings.HasPrefix(iface, "en")
		isUtun := strings.HasPrefix(iface, "utun")
		if !(isEn || isUtun) {
			t.Errorf("getIface() does not return expected interface: %s", iface)
		}
	}
}
