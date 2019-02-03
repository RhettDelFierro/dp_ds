package ds

import "strings"

type TrieNode struct {
	Key      string
	Parent   TrieNode
	Children map[string]TrieNode
	End      bool
}

func (tn *TrieNode) GetWord() string {
	var output []string
	node := tn

	for node != nil {
		output = append([]string{node.Key}, output...)
		node = &node.Parent
	}
	return strings.Join(output, "")
}
