package effects

import "bytes"

func Underline(text string) string {
	var ret bytes.Buffer

	for _, letter := range text {
		ret.WriteRune(letter)
		if letter != '\n' && letter != '\r' {
			ret.WriteRune(818)
		}
	}

	return ret.String()
}
