package effects

import "bytes"

var monospaceLowerChars = [...]rune{'𝚊', '𝚋', '𝚌', '𝚍', '𝚎', '𝚏', '𝚐', '𝚑', '𝚒', '𝚓', '𝚔', '𝚕', '𝚖', '𝚗', '𝚘', '𝚙', '𝚚', '𝚛', '𝚜', '𝚝', '𝚞', '𝚟', '𝚠', '𝚡', '𝚢', '𝚣', }
var monospaceUpperChars = [...]rune{'𝙰', '𝙱', '𝙲', '𝙳', '𝙴', '𝙵', '𝙶', '𝙷', '𝙸', '𝙹', '𝙺', '𝙻', '𝙼', '𝙽', '𝙾', '𝙿', '𝚀', '𝚁', '𝚂', '𝚃', '𝚄', '𝚅', '𝚆', '𝚇', '𝚈', '𝚉', }

func Monospace(text string) string {
	var ret bytes.Buffer

	for _, letter := range text {
		if letter >= 'a' && letter <= 'z' {
			ret.WriteString(string(monospaceLowerChars[letter-'a']))
		} else if letter >= 'A' && letter <= 'Z' {
			ret.WriteString(string(monospaceUpperChars[letter-'A']))
		} else {
			ret.WriteString(string(letter))
		}
	}

	return ret.String()
}
