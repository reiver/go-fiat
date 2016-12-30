package fiat

import (
	"strings"
	"unicode/utf8"
)

func extractFollowingLines(s string, endNeedle string) (string, error) {

	if "" == s {
		return "", nil
	}

	// Skip first line of text.
	for {
		r, n := utf8.DecodeRuneInString(s)
		if utf8.RuneError == r {
			return "", errRuneError
		}
		s = s[n:]
		if "" == s {
			return "", nil
		}

		if '\n' == r {
			break
		}
	}

	if "" == endNeedle {
		return s, nil
	}

	index := strings.Index(s, endNeedle)
	if -1 == index {
		return s, nil
	}
	if len(s) <= index {
		return s, nil
	}

	s = s[:index]

	return s, nil
}
