package main

import "bytes"

func underline(text string) string {
	var ret bytes.Buffer

	for _, letter := range text {
		ret.WriteString(string(letter))
		if letter != '\n' && letter != '\r' {
			ret.WriteString(string(818))
		}
	}

	return ret.String()
}
