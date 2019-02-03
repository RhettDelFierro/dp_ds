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

func TestContains(t *testing.T) {
	trie := NewTrie()
	trieNode := NewTrieNode()
	trie.Root = trieNode
	trie.Insert("blah")
	if trie.Contains("blah") == false {
		t.Errorf("Trie.Contains is not working. Wanted: '%v', got: '%v'", true, trie.Contains("blah"))
	}
}

func TestFind(t *testing.T) {
	trie := NewTrie()
	trieNode := NewTrieNode()
	trie.Root = trieNode
	trie.Insert("blah")
	trie.Insert("blahbar")
	if len(trie.Find("blah")) != 2 {
		t.Errorf("Trie.Find is not working. Wanted: '%d', got: '%d'", 2, len(trie.Find("blah")))
	}
}
