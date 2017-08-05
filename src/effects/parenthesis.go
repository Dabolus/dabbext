package effects

import "bytes"

func Parenthesis(text string) string {
	var ret bytes.Buffer

	for _, letter := range text {
		if letter >= 'a' && letter <= 'z' {
			ret.WriteString(string(letter + 9275))
		} else if letter >= 'A' && letter <= 'Z' {
			ret.WriteString(string(letter + 9307))
		} else if letter >= '1' && letter <= '9' {
			ret.WriteString(string(letter + 9283))
		} else {
			ret.WriteString(string(letter))
		}
	}

	return ret.String()
}
