package main

import "bytes"

var whiteSquaresChars = [...]rune{'🄰', '🄱', '🄲', '🄳', '🄴', '🄵', '🄶', '🄷', '🄸', '🄹', '🄺', '🄻', '🄼', '🄽', '🄾', '🄿', '🅀', '🅁', '🅂', '🅃', '🅄', '🅅', '🅆', '🅇', '🅈', '🅉', }

func whiteSquares(text string) string {
	var ret bytes.Buffer

	for _, letter := range text {
		if letter >= 'a' && letter <= 'z' {
			ret.WriteString(string(whiteSquaresChars[letter-'a']))
		} else if letter >= 'A' && letter <= 'Z' {
			ret.WriteString(string(whiteSquaresChars[letter-'A']))
		} else {
			ret.WriteString(string(letter))
		}
	}

	return ret.String()
}
