package main

import "bytes"

var risChars = [...]rune{'🇦', '🇧', '🇨', '🇩', '🇪', '🇫', '🇬', '🇭', '🇮', '🇯', '🇰', '🇱', '🇲', '🇳', '🇴', '🇵', '🇶', '🇷', '🇸', '🇹', '🇺', '🇻', '🇼', '🇽', '🇾', '🇿', }

func ris(text string) string {
	var ret bytes.Buffer

	for _, letter := range text {
		if letter >= 'a' && letter <= 'z' {
			ret.WriteRune(risChars[letter-'a'])
		} else if letter >= 'A' && letter <= 'Z' {
			ret.WriteRune(risChars[letter-'A'])
		} else {
			ret.WriteRune(letter)
		}
		ret.WriteRune(' ')
	}

	return ret.String()
}
