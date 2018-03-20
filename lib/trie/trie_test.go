package trie

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	trie := New()
	trie.Add("hello")
	trie.Add("world")

	assert.Equal(t, 2, len(trie.children))

	var ok bool
	var h *Trie

	h, ok = trie.children['h']
	assert.True(t, ok)
	assert.False(t, h.term)

	assert.Equal(t, 1, len(h.children))
	h, ok = h.children['e']
	assert.True(t, ok)
	assert.False(t, h.term)

	assert.Equal(t, 1, len(h.children))
	h, ok = h.children['l']
	assert.True(t, ok)
	assert.False(t, h.term)

	assert.Equal(t, 1, len(h.children))
	h, ok = h.children['l']
	assert.True(t, ok)
	assert.False(t, h.term)

	assert.Equal(t, 1, len(h.children))
	h, ok = h.children['o']
	assert.True(t, ok)

	assert.True(t, h.term)
	assert.Equal(t, 0, len(h.children))
}

var testStrings = []string{
	"pre",
	"prefix",
	"premium",
	"present",
	"prevail",
	"prevalent",
	"prevent",
	"proletariat",
	"prominent",
	"prop",
	"property",
	"proposal",
}

func TestPrefixSearch(t *testing.T) {
	trie := New()
	for _, s := range testStrings {
		trie.Add(s)
	}

	expected := testStrings[:7]
	actual := trie.PrefixSearch("pre")
	sort.Strings(actual)
	assert.Equal(t, expected, actual)

	expected = []string{"prevail", "prevalent", "prevent"}
	actual = trie.PrefixSearch("prev")
	sort.Strings(actual)
	assert.Equal(t, expected, actual)

	expected = testStrings[7:]
	actual = trie.PrefixSearch("pro")
	sort.Strings(actual)
	assert.Equal(t, expected, actual)

	expected = []string{}
	actual = trie.PrefixSearch("proz")
	sort.Strings(actual)
	assert.Equal(t, expected, actual)

	expected = testStrings
	actual = trie.PrefixSearch("")
	sort.Strings(actual)
	assert.Equal(t, expected, actual)
}

func TestRemove(t *testing.T) {
	trie := New()
	for _, s := range testStrings {
		trie.Add(s)
	}

	trie.Remove("pre")

	expected := testStrings[1:7]
	actual := trie.PrefixSearch("pre")
	sort.Strings(actual)
	assert.Equal(t, expected, actual)

	trie.Remove("prevent")

	expected = testStrings[1:6]
	actual = trie.PrefixSearch("pre")
	sort.Strings(actual)
	assert.Equal(t, expected, actual)

	trie.Remove("prop")

	expected = testStrings[10:]
	actual = trie.PrefixSearch("prop")
	sort.Strings(actual)
	assert.Equal(t, expected, actual)

	trie.Remove("prev")

	expected = testStrings[1:6]
	actual = trie.PrefixSearch("pre")
	sort.Strings(actual)
	assert.Equal(t, expected, actual, "Removing a non-existing word should do nothing")
}
