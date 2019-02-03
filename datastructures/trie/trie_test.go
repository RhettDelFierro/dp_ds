package ds

import (
	"testing"
)

func NewTrieNode() *TrieNode {
	return &TrieNode{}
}

func NewTrie() *Trie {
	return &Trie{}
}

func TestInsert(t *testing.T) {
	trie := NewTrie()
	trieNode := NewTrieNode()
	trie.Root = trieNode

	trie.Insert("blah")
	if len(trie.Root.Children["b"].GetWord()) == 0 {
		t.Errorf("Trie.Insert and TrieNode.GetWord do not work together. Wanted: '%d', got: '%d'", 4, len(trie.Root.Children["b"].GetWord()))
	}
}
