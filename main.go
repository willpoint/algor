package main

import (
	"fmt"

	"github.com/willpoint/algor/pkg/ds"
)

func main() {

	ss := []string{"bull", "adam", "ope", "busayo"}

	trie := ds.NewTrie()

	for _, j := range ss {
		if err := trie.Insert(j); err != nil {
			fmt.Println(err)
			continue
		}
	}
	var f func(map[rune]*ds.TrieNode)
	f = func(tn map[rune]*ds.TrieNode) {
		for i, j := range tn {
			if j == nil {
				return
			}
			fmt.Printf("%s", string(i))
			f(j.Children)
			fmt.Println()
		}
	}
	f(trie.Root.Children)
	for _, j := range ss {
		fmt.Println(trie.Search(j))
	}

}
