package trie

type Trie struct {
	val      rune
	depth    int
	term     bool
	parent   *Trie
	children map[rune]*Trie
}

func New() *Trie {
	return &Trie{parent: nil, depth: 0, children: make(map[rune]*Trie)}
}

// Add a new word to the trie.
func (t *Trie) Add(s string) {
	node := t
	for _, r := range s {
		n, ok := node.children[r]
		if !ok {
			n = &Trie{
				val:      r,
				depth:    node.depth + 1,
				term:     false,
				parent:   node,
				children: make(map[rune]*Trie),
			}
			node.children[r] = n
		}
		node = n
	}
	node.term = true
}

// Remove the specified word from the trie.
// If the given string is not a full word in the trie
// (even if it is a prefix of other words), do nothing.
func (t *Trie) Remove(s string) {
	node := t.findNode([]rune(s))
	if node == nil || !node.term {
		return
	}

	node.term = false
	for n := node.parent; n != nil && len(node.children) == 0; node, n = n, n.parent {
		delete(n.children, node.val)
	}
}

// Return a slice of all words in the trie which begin with the specified prefix.
func (t *Trie) PrefixSearch(prefix string) []string {
	node := t.findNode([]rune(prefix))
	if node == nil {
		return make([]string, 0)
	}
	return node.collect()
}

func (t *Trie) findNode(runes []rune) *Trie {
	for _, r := range runes {
		if n, ok := t.children[r]; ok {
			t = n
		} else {
			return nil
		}
	}

	return t
}

// Return a slice of all words in the tree rooted at the target node.
func (t *Trie) collect() []string {
	ret := make([]string, 0)

	nodes := []*Trie{t}
	for l := len(nodes); l != 0; l = len(nodes) {
		i := l - 1
		n := nodes[i]
		nodes = nodes[:i]
		for _, c := range n.children {
			nodes = append(nodes, c)
		}
		if n.term {
			ret = append(ret, n.word())
		}
	}

	return ret
}

func (t *Trie) word() string {
	runes := make([]rune, t.depth)

	for ; t.depth != 0; t = t.parent {
		runes[t.depth - 1] = t.val
	}

	return string(runes)
}
