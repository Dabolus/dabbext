package effects

import "bytes"

func Vaporwave(text string) string {
	var ret bytes.Buffer

	// Loop through each letter
	for _, letter := range text {
		// Normal letters range
		if letter > 32 && letter < 127 {
			ret.WriteString(string(letter + 65248))
		// Space
		} else if letter == ' ' {
			ret.WriteString("　")
		// Special parenthesis
		} else if letter == '⦅' || letter == '⦆' {
			ret.WriteString(string(letter + 54746))
		// No full-width corresponding letter, just add the normal one
		} else {
			ret.WriteString(string(letter))
		}
	}

	return ret.String()
}
