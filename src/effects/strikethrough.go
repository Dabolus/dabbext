package effects

import "bytes"

func Strikethrough(text string) string {
	var ret bytes.Buffer

	for _, letter := range text {
		ret.WriteString(string(letter))
		if letter != '\n' && letter != '\r' {
			ret.WriteString(string(822))
		}
	}

	return ret.String()
}