package main

import "bytes"

func overline(text string) string {
	var ret bytes.Buffer

	for _, letter := range text {
		ret.WriteString(string(letter))
		if letter != '\n' && letter != '\r' {
			ret.WriteString(string(773))
		}
	}

	return ret.String()
}
