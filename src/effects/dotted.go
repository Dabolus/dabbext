package effects

import "bytes"

func Dotted(text string) string {
	var ret bytes.Buffer

	for _, letter := range text {
		ret.WriteRune(letter)
		if letter != '\n' && letter != '\r' {
			ret.WriteRune(775)
		}
	}

	return ret.String()
}
