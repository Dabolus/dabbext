package effects

import "bytes"

var blackCirclesChars = [...]rune{'🅐', '🅑', '🅒', '🅓', '🅔', '🅕', '🅖', '🅗', '🅘', '🅙', '🅚', '🅛', '🅜', '🅝', '🅞', '🅟', '🅠', '🅡', '🅢', '🅣', '🅤', '🅥', '🅦', '🅧', '🅨', '🅩', }

func BlackCircles(text string) string {
	var ret bytes.Buffer

	for _, letter := range text {
		if letter >= 'a' && letter <= 'z' {
			ret.WriteString(string(blackCirclesChars[letter - 'a']))
		} else if letter >= 'A' && letter <= 'Z' {
			ret.WriteString(string(blackCirclesChars[letter - 'A']))
		} else if letter >= '1' && letter <= '9' {
			ret.WriteString(string(letter + 10073))
		} else if letter == '0' {
			ret.WriteString("⓿")
		} else {
			ret.WriteString(string(letter))
		}
	}

	return ret.String()
}
