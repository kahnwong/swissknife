package generate

import (
	"reflect"
	"testing"
)

func Test_generateQRCode(t *testing.T) {
	qrcode, err := generateQRCode("https://example.com")
	if err != nil {
		t.Fatal(err)
	} else {
		if reflect.TypeOf(qrcode).Kind() != reflect.String {
			t.Errorf("generateQRCode(\"https://example.com\") = %v, want string", qrcode)
		}
	}
}
