package main

import "bytes"

func whiteCircles(text string) string {
	var ret bytes.Buffer

	for _, letter := range text {
		if letter >= 'a' && letter <= 'z' {
			ret.WriteRune(letter + 9327)
		} else if letter >= 'A' && letter <= 'Z' {
			ret.WriteRune(letter + 9333)
		} else if letter >= '1' && letter <= '9' {
			ret.WriteRune(letter + 10063)
		} else if letter == '0' {
			ret.WriteRune('⓪')
		} else {
			ret.WriteRune(letter)
		}
	}

	return ret.String()
}
