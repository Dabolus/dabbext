package main

import "bytes"

func dotted(text string) string {
	var ret bytes.Buffer

	for _, letter := range text {
		ret.WriteString(string(letter))
		if letter != '\n' && letter != '\r' {
			ret.WriteString(string(775))
		}
	}

	return ret.String()
}
