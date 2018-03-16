package tk

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"timekeeper/lib/tkerr"
)

func TestCanonicalizeUsername(t *testing.T) {
	var str string
	var err error

	str, err = canonicalizeUsername("hello")
	assert.Nil(t, err)
	assert.Equal(t, "hello", str, "`hello' is already in canonical form")

	str, err = canonicalizeUsername("HeLLo123")
	assert.Nil(t, err)
	assert.Equal(t, "hello123", str, "uppercase letters should be lowercased")

	str, err = canonicalizeUsername("Hello World")
	assert.Nil(t, err)
	assert.Equal(t, "hello_world", str, "spaces should be converted to underscores")

	str, err = canonicalizeUsername("")
	if assert.NotNil(t, err, "usernames cannot be empty") {
		if assert.IsType(t, &tkerr.TKError{}, err) {
			e, _ := err.(*tkerr.TKError)
			assert.Equal(t, tkerr.InvalidUsername, e.Code)
		}
	}

	str, err = canonicalizeUsername("abcdefghijklm")
	if assert.NotNil(t, err, "usernames have a max length of 12") {
		if assert.IsType(t, &tkerr.TKError{}, err) {
			e, _ := err.(*tkerr.TKError)
			assert.Equal(t, tkerr.InvalidUsername, e.Code)
		}
	}

	str, err = canonicalizeUsername("a28_i#")
	if assert.NotNil(t, err, "usernames cannot contain special characters") {
		if assert.IsType(t, &tkerr.TKError{}, err) {
			e, _ := err.(*tkerr.TKError)
			assert.Equal(t, tkerr.InvalidUsername, e.Code)
		}
	}
}
