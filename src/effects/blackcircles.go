package effects

import "bytes"

var blackCirclesChars = [...]rune{'🅐', '🅑', '🅒', '🅓', '🅔', '🅕', '🅖', '🅗', '🅘', '🅙', '🅚', '🅛', '🅜', '🅝', '🅞', '🅟', '🅠', '🅡', '🅢', '🅣', '🅤', '🅥', '🅦', '🅧', '🅨', '🅩', }

func BlackCircles(text string) string {
	var ret bytes.Buffer

	for _, letter := range text {
		if letter >= 'a' && letter <= 'z' {
			ret.WriteRune(blackCirclesChars[letter - 'a'])
		} else if letter >= 'A' && letter <= 'Z' {
			ret.WriteRune(blackCirclesChars[letter - 'A'])
		} else if letter >= '1' && letter <= '9' {
			ret.WriteRune(letter + 10073)
		} else if letter == '0' {
			ret.WriteRune('⓿')
		} else {
			ret.WriteRune(letter)
		}
	}

	return ret.String()
}
