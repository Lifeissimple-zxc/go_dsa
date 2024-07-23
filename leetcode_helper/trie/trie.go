package trie

import (
	"strings"
)

const alphabetSize = 26

// Node represents a trie node
type Node struct {
	children [alphabetSize]*Node
	isEnd    bool
}

// Trie represents a Trie
type Trie struct {
	root *Node
}

// New trie returns a new trie
func NewTrie() *Trie {
	return &Trie{root: &Node{}}
}

// Insert adds a new word to the trie
func (t *Trie) Insert(w string) {
	wLen := len(w)
	cursor := t.root
	for i := 0; i < wLen; i++ {
		// offsetting with 'a' based on ascii tables
		charIx := (strings.ToLower(w)[i]) - 'a'
		// empty char --> time to append
		if cursor.children[charIx] == nil {
			cursor.children[charIx] = &Node{}
		}
		// update cursor for traversal
		cursor = cursor.children[charIx]
	}
	// mark end of the word
	cursor.isEnd = true

}

// Search checks if a word is in the trie
func (t *Trie) Search(w string) bool {
	wLen := len(w)
	cursor := t.root
	for i := 0; i < wLen; i++ {
		// offsetting with 'a' based on ascii tables
		charIx := (strings.ToLower(w)[i]) - 'a'
		// empty char --> time to append
		if cursor.children[charIx] == nil {
			return false
		}
		// update cursor for traversal
		cursor = cursor.children[charIx]
	}
	return cursor.isEnd
}
