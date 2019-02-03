package ds

import (
	"strings"
)

type TrieNode struct {
	Key      string
	Parent   *TrieNode
	Children map[string]*TrieNode
	End      bool
}

func (tn *TrieNode) GetWord() string {
	var output []string
	node := tn

	for node != nil {
		output = append([]string{node.Key}, output...)
		node = node.Parent
	}
	return strings.Join(output, "")
}

type Trie struct {
	Root *TrieNode
}

func (t *Trie) Insert(word string) {
	node := t.Root
	for i := 0; i < len(word); i++ {
		key := string([]rune(word)[i])

		if node.Children[key] == nil {
			node.Children = make(map[string]*TrieNode)
			node.Children[key] = &TrieNode{Key: key}
			node.Children[key].Parent = node
		}

		node = node.Children[key]
		if i == len(word)-1 {
			node.End = true
		}
	}
}
