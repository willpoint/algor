package trie

import (
	"errors"
)

var (
	// ErrPrefixStringExists ...
	ErrPrefixStringExists = errors.New("prefix string exists for this new string")
)

// Trie provides a string searching algorithm
// that preprocesses texts to enable quick search.
// The root is the start of a trie which records every
// character of a string in a collection such that
// from the root of the tree to any external node within
// the trie - the concatenation of labels give a string
// which is a member of string collection
// An improved trie saves more space by saving strings instead of
// characters and is essentially a Trie with more constraints
// referred to as a patricia trie. It ensures that each internal
// node in the trie has at least two children, by compressing chains
// of single-child nodes into individual edges and avoiding redundant
// nodes within the Trie
// eg. if vi is redundant for i = 1, ..., k - 1
// then v0 and vk are not redundant prevented by a concatenation
// of the labels vi, ..., vk
type Trie struct {
	Root  *TrieNode
	Words int
}

// NewTrie returns a Trie with a root having initialized Children
func NewTrie() *Trie {
	root := &TrieNode{
		Children: make(map[rune]*TrieNode),
	}
	return &Trie{root, 0}
}

// TrieNode uses a map to achieve a
// constant time search 0(1) or 0(log d)
// The External property serves as a sentinel
// to ensure no new string to be inserted has a prefix
// of an existing string
type TrieNode struct {
	Label    rune
	Children map[rune]*TrieNode
	External bool
}

// Insert uses an incremental approach that inserts the string s
// one at a time by first tracing the path associated with s in t.
func (t *Trie) Insert(s string) error {

	currNode := t.Root
	index := 0

	for _, j := range s {
		if _, ok := currNode.Children[j]; ok {
			currNode = currNode.Children[j]
			index++
		}
	}

	if currNode.External {
		return ErrPrefixStringExists
	}

	for i, j := range s[index:] {
		tnode := &TrieNode{
			Label:    j,
			Children: make(map[rune]*TrieNode),
		}
		currNode.Children[j] = tnode
		currNode = tnode
		if i == len(s[index:])-1 {
			currNode.External = true
			t.Words++
		}
	}
	return nil
}

// Search walks through the Trie T from the root
// and checks for every node if i character of the string s
// matches paths from the root to an external node
// if after the walk it terminates with the current node v
// empty then the string does not exist and returns a false
// it returns true otherwise
func (t *Trie) Search(s string) bool {
	currNode := t.Root
	for _, j := range s {
		if _, ok := currNode.Children[j]; ok {
			currNode = currNode.Children[j]
		}
	}
	if currNode.External {
		return true
	}
	return false
}
