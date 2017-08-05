package effects

import "bytes"

var blackSquaresChars = [...]rune{'🅰', '🅱', '🅲', '🅳', '🅴', '🅵', '🅶', '🅷', '🅸', '🅹', '🅺', '🅻', '🅼', '🅽', '🅾', '🅿', '🆀', '🆁', '🆂', '🆃', '🆄', '🆅', '🆆', '🆇', '🆈', '🆉', }

func BlackSquares(text string) string {
	var ret bytes.Buffer

	for _, letter := range text {
		if letter >= 'a' && letter <= 'z' {
			ret.WriteString(string(blackSquaresChars[letter-'a']))
		} else if letter >= 'A' && letter <= 'Z' {
			ret.WriteString(string(blackSquaresChars[letter-'A']))
		} else {
			ret.WriteString(string(letter))
		}
	}

	return ret.String()
}
