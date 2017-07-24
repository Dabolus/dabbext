package main

import "bytes"

func whiteCircles(text string) string {
	var ret bytes.Buffer

	for _, letter := range text {
		if letter >= 'a' && letter <= 'z' {
			ret.WriteString(string(letter + 9327))
		} else if letter >= 'A' && letter <= 'Z' {
			ret.WriteString(string(letter + 9333))
		} else if letter >= '1' && letter <= '9' {
			ret.WriteString(string(letter + 10063))
		} else if letter == '0' {
			ret.WriteString("⓪")
		} else {
			ret.WriteString(string(letter))
		}
	}

	return ret.String()
}
