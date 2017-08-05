package effects

import "bytes"

var ideographChars = [...]rune{'丹', '乃', '匚', '刀', 'モ', '下', 'ム', '卄', '工', 'Ｊ', 'Ｋ', 'ㄥ', '爪', 'れ', '口', 'ㄗ', 'Ｑ', '尺', 'ち', '匕', 'Ｕ', 'Ｖ', '山', 'メ', 'ㄚ', '乙', }

func Ideographs(text string) string {
	var ret bytes.Buffer

	for _, letter := range text {
		if letter >= 'a' && letter <= 'z' {
			ret.WriteRune(ideographChars[letter - 'a'])
		} else if letter >= 'A' && letter <= 'Z' {
			ret.WriteRune(ideographChars[letter - 'A'])
		} else {
			ret.WriteRune(letter)
		}
	}

	return ret.String()
}
