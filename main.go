package main

import "fmt"

func main() {

	// ss := []string{"bear", "bell", "bid", "bull", "bully", "buy", "sell", "stock", "stop"}

	// trie := ds.NewTrie()

	// for _, j := range ss {
	// 	if err := trie.Insert(j); err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	}
	// }
	// var f func(map[rune]*ds.TrieNode)
	// f = func(tn map[rune]*ds.TrieNode) {
	// 	for i, j := range tn {
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
	a := "man"
	b := "woman"
	i := string(a[0])
	j := string(b[0])
	fmt.Println(i + j)

}
