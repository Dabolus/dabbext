package effects

import (
	"bytes"
	"math/rand"
)

var (
	zalgoUp = [...]rune{
		'\u030d', /* ̍ */ '\u030e', /* ̎ */ '\u0304', /* ̄ */ '\u0305', /* ̅ */
		'\u033f', /* ̿ */ '\u0311', /* ̑ */ '\u0306', /* ̆ */ '\u0310', /* ̐ */
		'\u0352', /* ͒ */ '\u0357', /* ͗ */ '\u0351', /* ͑ */ '\u0307', /* ̇  */
		'\u0308', /* ̈ */ '\u030a', /* ̊ */ '\u0342', /* ͂ */ '\u0343', /* ̓ */
		'\u0344', /* ̈́ */ '\u034a', /* ͊ */ '\u034b', /* ͋ */ '\u034c', /* ͌ */
		'\u0303', /* ̃  */ '\u0302', /* ̂ */ '\u030c', /* ̌ */ '\u0350', /* ͐ */
		'\u0300', /* ̀  */ '\u0301', /* ́  */ '\u030b', /* ̋ */ '\u030f', /* ̏ */
		'\u0312', /* ̒ */ '\u0313', /* ̓ */ '\u0314', /* ̔ */ '\u033d', /* ̽ */
		'\u0309', /* ̉  */ '\u0363', /* ͣ */ '\u0364', /* ͤ */ '\u0365', /* ͥ */
		'\u0366', /* ͦ */ '\u0367', /* ͧ */ '\u0368', /* ͨ */ '\u0369', /* ͩ */
		'\u036a', /* ͪ */ '\u036b', /* ͫ */ '\u036c', /* ͬ */ '\u036d', /* ͭ */
		'\u036e', /* ͮ */ '\u036f', /* ͯ */ '\u033e', /* ̾ */ '\u035b', /* ͛ */
		'\u0346', /* ͆ */ '\u031a', /* ̚ */
	}
	zalgoMid = [...]rune{
		'\u0315', /* ̕ */ '\u031b', /* ̛ */ '\u0340', /* ̀  */ '\u0341', /* ́  */
		'\u0358', /* ͘ */ '\u0321', /* ̡ */ '\u0322', /* ̢ */ '\u0327', /* ̧ */
		'\u0328', /* ̨ */ '\u0334', /* ̴ */ '\u0335', /* ̵ */ '\u0336', /* ̶ */
		'\u034f', /* ͏  */ '\u035c', /* ͜ */ '\u035d', /* ͝ */ '\u035e', /* ͞ */
		'\u035f', /* ͟ */ '\u0360', /* ͠ */ '\u0362', /* ͢ */ '\u0338', /* ̸ */
		'\u0337', /* ̷ */ '\u0361', /* ͡ */ '\u0489', /* ҉_*/
	}
	zalgoDown = [...]rune{
		'\u0316', /* ̖  */ '\u0317', /* ̗  */ '\u0318', /* ̘  */ '\u0319', /* ̙  */
		'\u031c', /* ̜  */ '\u031d', /* ̝  */ '\u031e', /* ̞  */ '\u031f', /* ̟  */
		'\u0320', /* ̠  */ '\u0324', /* ̤  */ '\u0325', /* ̥  */ '\u0326', /* ̦  */
		'\u0329', /* ̩  */ '\u032a', /* ̪  */ '\u032b', /* ̫  */ '\u032c', /* ̬  */
		'\u032d', /* ̭  */ '\u032e', /* ̮  */ '\u032f', /* ̯  */ '\u0330', /* ̰  */
		'\u0331', /* ̱  */ '\u0332', /* ̲  */ '\u0333', /* ̳  */ '\u0339', /* ̹  */
		'\u033a', /* ̺  */ '\u033b', /* ̻  */ '\u033c', /* ̼  */ '\u0345', /* ͅ  */
		'\u0347', /* ͇  */ '\u0348', /* ͈  */ '\u0349', /* ͉  */ '\u034d', /* ͍  */
		'\u034e', /* ͎  */ '\u0353', /* ͓  */ '\u0354', /* ͔  */ '\u0355', /* ͕  */
		'\u0356', /* ͖  */ '\u0359', /* ͙  */ '\u035a', /* ͚  */ '\u0323', /* ̣  */
	}
)


func isZalgoChar(char rune) bool {
	for _, zalgoChar := range zalgoUp {
		if char == zalgoChar {
			return true
		}
	}
	for _, zalgoChar := range zalgoMid {
		if char == zalgoChar {
			return true
		}
	}
	for _, zalgoChar := range zalgoDown {
		if char == zalgoChar {
			return true
		}
	}
	return false
}


func Zalgo(text string) string {
	var ret bytes.Buffer

	for _, letter := range text {
		// We already have a zalgo character, nothing to do here
		if isZalgoChar(letter) == true {
			continue
		}

		// Add the normal character
		ret.WriteString(string(letter))

		numUp, numMid, numDown := rand.Intn(16) / 2 + 1, rand.Intn(6) / 2, rand.Intn(16) / 2 + 1

		for i := 0; i < numUp; i++ {
			ret.WriteString(string(zalgoUp[rand.Intn(len(zalgoUp))]))
		}
		for i := 0; i < numMid; i++ {
			ret.WriteString(string(zalgoMid[rand.Intn(len(zalgoMid))]))
		}
		for i := 0; i < numDown; i++ {
			ret.WriteString(string(zalgoDown[rand.Intn(len(zalgoDown))]))
		}
	}

	return ret.String()
}