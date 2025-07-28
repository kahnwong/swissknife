package generate

import (
	"github.com/kahnwong/swissknife/internal/generate"
	"github.com/spf13/cobra"
)

var generateQRCodeCmd = &cobra.Command{
	Use:   "qrcode",
	Short: "Generate QR code",
	Long:  `Generate QR code from URL (either as an arg or from clipboard) and copy resulting image to clipboard.`,
	Run: func(cmd *cobra.Command, args []string) {
		generate.QRCode(args)
	},
}

func init() {
	Cmd.AddCommand(generateQRCodeCmd)
}
