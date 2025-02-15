package ext

import (
	"regexp"
)

var regexesInitialized = false
var urlRegex *regexp.Regexp

func initializeRegexes() error {
	if regexesInitialized {
		return nil
	}
	var err error

	//goland:noinspection RegExpRedundantEscape,RegExpDuplicateCharacterInClass,GoConvertStringLiterals
	urlRegex, err = regexp.Compile(`[(http(s)?):\/\/(www\.)?a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`)
	if err != nil {
		return err
	}

	regexesInitialized = true
	return nil
}

func IsValidUrl(val string) (match bool, err error) {
	err = initializeRegexes()
	if err != nil {
		return false, err
	}

	return urlRegex.Match([]byte(val)), nil
}
