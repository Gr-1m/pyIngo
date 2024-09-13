package ascii

import "strings"

const (
	Digits          = "0123456789"
	OctDigits       = "01234567"
	HexDigits       = "0123456789abcdefABCDEF"
	Ascii_lowercase = "abcdefghijklmnopqrstuvwxyz"
	Ascii_uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Ascii_letters   = Ascii_lowercase + Ascii_uppercase

	WhiteSpace   = " \t\n\r\x0b\x0c"
	Punctuation  = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	Ascii_symbol = Punctuation + WhiteSpace
	Printable    = Ascii_letters + Ascii_symbol

	UUIDstr = HexDigits + "-_"
	FLAGstr = HexDigits + "-_}xX"
)

func Capwords(s string) string {
	return strings.ToUpper(s)
}
