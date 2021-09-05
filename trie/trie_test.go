package trie

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := Constructor()

	trie.Insert("how")
	trie.Insert("hi")
	trie.Insert("her")
	trie.Insert("hello")
	trie.Insert("so")
	trie.Insert("see")

	fmt.Println(trie.Search("hi"))
	fmt.Println(trie.Search("he"))
	fmt.Println(trie.Search("hello"))
	fmt.Println(trie.StartsWith("he"))
	fmt.Println(trie.StartsWith("see"))

	for _, v := range trie.GetPrefixList("s") {
		fmt.Println(v)
	}

	for _, v := range trie.GetPrefixList("he") {
		fmt.Println(v)
	}
}
