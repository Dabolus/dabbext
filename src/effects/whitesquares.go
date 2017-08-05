package effects

import "bytes"

var whiteSquaresChars = [...]rune{'🄰', '🄱', '🄲', '🄳', '🄴', '🄵', '🄶', '🄷', '🄸', '🄹', '🄺', '🄻', '🄼', '🄽', '🄾', '🄿', '🅀', '🅁', '🅂', '🅃', '🅄', '🅅', '🅆', '🅇', '🅈', '🅉', }

func WhiteSquares(text string) string {
	var ret bytes.Buffer

	for _, letter := range text {
		if letter >= 'a' && letter <= 'z' {
			ret.WriteRune(whiteSquaresChars[letter-'a'])
		} else if letter >= 'A' && letter <= 'Z' {
			ret.WriteRune(whiteSquaresChars[letter-'A'])
		} else {
			ret.WriteRune(letter)
		}
	}

	return ret.String()
}
