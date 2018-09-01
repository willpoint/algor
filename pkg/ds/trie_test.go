package ds

import (
	"testing"
)

func TestInsert(t *testing.T) {

	trie := NewTrie()
	ss := []string{"one", "two", "exam", "example"}
	words := len(ss) - 1
	for _, s := range ss {
		if err := trie.Insert(s); err != nil {
			continue
		}
	}
	if trie.Words != words {
		t.Errorf("expect word count to equal %d, got %d", words, trie.Words)
	}
	err := trie.Insert("ones")
	if err != ErrPrefixStringExists {
		t.Errorf("expect error to be non nil")
	}

	// var f func(map[rune]*TrieNode)
	// f = func(tn map[rune]*TrieNode) {
	// 	for i, j := range trie {
	// 		if j == nil {
	// 			return
	// 		}
	// 		fmt.Printf("%s", string(i))
	// 		if j.External {
	// 			fmt.Println("|")
	// 		}
	// 		f(j.Children)
	// 	}
	// }
	// f(trie.Root.Children)

}
