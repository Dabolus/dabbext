package main

import "bytes"

func denied(text string) string {
	var ret bytes.Buffer

	for _, letter := range text {
		ret.WriteString(string(letter))
		if letter != '\n' && letter != '\r' && letter != '\t' && letter != ' ' {
			ret.WriteString(string(8416))
		}
	}

	return ret.String()
}
