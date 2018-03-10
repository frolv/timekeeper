package tk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidUsername(t *testing.T) {
	var b bool

	b = validUsername("hello")
	assert.Equal(t, true, b, "username `hello' should be valid")

	b = validUsername("")
	assert.Equal(t, false, b, "usernames cannot be empty")

	b = validUsername("abcdefghijklm")
	assert.Equal(t, false, b, "usernames have a max length of 12")

	b = validUsername("a28_i#")
	assert.Equal(t, false, b, "usernames cannot contain special characters")
}
