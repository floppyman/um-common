package ext

import (
	"regexp"
	"strings"
	"unicode"
)

const ansiColorStripRegexStr = "\x1B(?:[@-Z\\-_]|\\[[0-?]*[ -/]*[@-~])"

var ansiColorStripRegex = regexp.MustCompile(ansiColorStripRegexStr)

// Cut cuts a string to the specified length, if it is longer than the specified length
//
//goland:noinspection GoUnusedExportedFunction
func Cut(val string, length int) string {
	if len(val) >= length {
		return val[0:length]
	}
	return val
}

// CutUnicode returns a new string at the specified length, ensuring that unicode characters are not cut in the middle resulting in broken characters, or the original string if it is shorter than length
//
//goland:noinspection GoUnusedExportedFunction
func CutUnicode(str string, length int) string {
	count := 0
	text := strings.Map(func(r rune) rune {
		if count == length {
			return -1
		}
		if unicode.IsPrint(r) {
			count++
			return r
		}
		return -1
	}, str)
	return text
}

// StripAnsiColors strips all ansi color values from a string
//
//goland:noinspection GoUnusedExportedFunction
func StripAnsiColors(str string) string {
	return ansiColorStripRegex.ReplaceAllString(str, "")
}
