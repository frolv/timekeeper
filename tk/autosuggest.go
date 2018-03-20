package tk

import (
	"timekeeper/lib/trie"
	"sort"
)

var names *trie.Trie
var displayNames map[string]string

func autoSuggestInit() {
	var accounts []Account

	db.Select("username, display_name").Find(&accounts)

	names = trie.New()
	displayNames = make(map[string]string)

	for _, acc := range accounts {
		names.Add(acc.Username)
		displayNames[acc.Username] = acc.DisplayName
	}
}

// Return the first `num` usernames which start with the given query string.
func AutoSuggest(query string, num int) []string {
	usernames := names.PrefixSearch(query)
	sort.Strings(usernames)

	if num > len(usernames) {
		num = len(usernames)
	}
	usernames = usernames[:num]

	for i, name := range usernames {
		usernames[i] = displayNames[name]
	}

	return usernames
}
