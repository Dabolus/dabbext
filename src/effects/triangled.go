package effects

import "bytes"

func Triangled(text string) string {
	var ret bytes.Buffer

	for _, letter := range text {
		ret.WriteRune(letter)
		if letter != '\n' && letter != '\r' {
			ret.WriteRune(8420)
		}
	}

	return ret.String()
}
