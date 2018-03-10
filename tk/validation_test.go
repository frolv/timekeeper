package tk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidUsername(t *testing.T) {
	var b bool

	b = validUsername("hello")
	assert.Equal(t, b, true, "username `hello' should be valid")

	b = validUsername("")
	assert.Equal(t, b, false, "usernames cannot be empty")

	b = validUsername("abcdefghijklm")
	assert.Equal(t, b, false, "usernames have a max length of 12")

	b = validUsername("a28_i#")
	assert.Equal(t, b, false, "usernames cannot contain special characters")
}
