package effects

import "bytes"

var frakturLowerChars = [...]rune{'𝖆', '𝖇', '𝖈', '𝖉', '𝖊', '𝖋', '𝖌', '𝖍', '𝖎', '𝖏', '𝖐', '𝖑', '𝖒', '𝖓', '𝖔', '𝖕', '𝖖', '𝖗', '𝖘', '𝖙', '𝖚', '𝖛', '𝖜', '𝖝', '𝖞', '𝖟' }
var frakturUpperChars = [...]rune{'𝕬', '𝕭', '𝕮', '𝕯', '𝕰', '𝕱', '𝕲', '𝕳', '𝕴', '𝕵', '𝕶', '𝕷', '𝕸', '𝕹', '𝕺', '𝕻', '𝕼', '𝕽', '𝕾', '𝕿', '𝖀', '𝖁', '𝖂', '𝖃', '𝖄', '𝖅', }

func Fraktur(text string) string {
	var ret bytes.Buffer

	for _, letter := range text {
		if letter >= 'a' && letter <= 'z' {
			ret.WriteRune(frakturLowerChars[letter-'a'])
		} else if letter >= 'A' && letter <= 'Z' {
			ret.WriteRune(frakturUpperChars[letter-'A'])
		} else {
			ret.WriteRune(letter)
		}
	}

	return ret.String()
}
