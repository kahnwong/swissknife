package generate

import (
	"testing"
)

func TestGenerateQRCode(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{
			name:    "valid URL",
			url:     "https://example.com",
			wantErr: false,
		},
		{
			name:    "empty URL",
			url:     "",
			wantErr: true, // QR code library doesn't accept empty strings
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			png, stdout, err := generateQRCode(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateQRCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if len(png) == 0 {
					t.Error("generateQRCode() returned empty PNG")
				}
				if len(stdout) == 0 {
					t.Error("generateQRCode() returned empty stdout")
				}
			}
		})
	}
}
