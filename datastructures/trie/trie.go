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
			//https://stackoverflow.com/questions/27267900/runtime-error-assignment-to-entry-in-nil-map
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

func (t *Trie) Contains(word string) bool {
	node := t.Root
	for i := 0; i < len(word); i++ {
		key := string([]rune(word)[i])
		if node.Children[key] != nil {
			node = node.Children[key]
		} else {
			return false
		}
	}
	return node.End
}

func (t *Trie) Find(prefix string) []string {
	node := t.Root
	var output []string

	for i := 0; i < len(prefix); i++ {
		key := string([]rune(prefix)[i])
		if node.Children[key] != nil {
			node = node.Children[key]
		} else {
			return output
		}
	}

	FindAllWords(node, &output)

	return output
}

func FindAllWords(node *TrieNode, arr *[]string) {
	if node.End {
		*arr = append([]string{node.GetWord()}, *arr...)
	}

	for key := range node.Children {
		FindAllWords(node.Children[key], arr)
	}
}
