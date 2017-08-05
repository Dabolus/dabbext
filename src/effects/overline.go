package effects

import "bytes"

func Overline(text string) string {
	var ret bytes.Buffer

	for _, letter := range text {
		ret.WriteRune(letter)
		if letter != '\n' && letter != '\r' {
			ret.WriteRune(773)
		}
	}

	return ret.String()
}
