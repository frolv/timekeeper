package tk

import "unicode"

func validUsername(username string) bool {
	valid := true
	for _, c := range username {
		if !unicode.IsDigit(c) && !unicode.IsLetter(c) && c != '_' {
			valid = false
			break
		}
	}

	return valid && len(username) > 0 && len(username) <= 12
}
