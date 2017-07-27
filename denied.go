package main

import "bytes"

func denied(text string) string {
	var ret bytes.Buffer

	for _, letter := range text {
		ret.WriteRune(letter)
		if letter != '\n' && letter != '\r' && letter != '\t' && letter != ' ' {
			ret.WriteRune(8416)
		}
	}

	return ret.String()
}
