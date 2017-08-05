package effects

import "bytes"

func Parenthesis(text string) string {
	var ret bytes.Buffer

	for _, letter := range text {
		if letter >= 'a' && letter <= 'z' {
			ret.WriteRune(letter + 9275)
		} else if letter >= 'A' && letter <= 'Z' {
			ret.WriteRune(letter + 9307)
		} else if letter >= '1' && letter <= '9' {
			ret.WriteRune(letter + 9283)
		} else {
			ret.WriteRune(letter)
		}
	}

	return ret.String()
}
