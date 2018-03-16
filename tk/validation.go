package tk

import (
	"bytes"
	"unicode"

	"timekeeper/lib/tkerr"
)

func canonicalizeUsername(username string) (string, error) {
	var buf bytes.Buffer

	if len(username) == 0 || len(username) > 12 {
		return "", tkerr.Create(tkerr.InvalidUsername)
	}

	for _, c := range username {
		switch {
		case '0' <= c && c <= '9', 'a' <= c && c <= 'z', c == '_':
			buf.WriteRune(c)
		case 'A' <= c && c <= 'Z':
			buf.WriteRune(unicode.ToLower(c))
		case c == ' ':
			buf.WriteRune('_')
		default:
			return "", tkerr.Create(tkerr.InvalidUsername)
		}
	}

	return buf.String(), nil
}
