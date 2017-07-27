package main

import "bytes"

func strikethrough(text string) string {
	var ret bytes.Buffer

	for _, letter := range text {
		ret.WriteRune(letter)
		if letter != '\n' && letter != '\r' {
			ret.WriteRune(822)
		}
	}

	return ret.String()
}
