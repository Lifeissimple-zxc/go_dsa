package trie

import (
	"strings"
)

const (
	alphabetSize = 26
	offsetASCIIa = 97
)

// Node represents a trie node
type Node struct {
	children [alphabetSize]*Node
	isEnd    bool
}

// GetChildren is a getter for children
func (n *Node) GetChildren() [alphabetSize]*Node {
	return n.children
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
		charIx := (strings.ToLower(w)[i]) - offsetASCIIa
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
		charIx := (strings.ToLower(w)[i]) - offsetASCIIa
		// empty char --> time to append
		if cursor.children[charIx] == nil {
			return false
		}
		// update cursor for traversal
		cursor = cursor.children[charIx]
	}
	return cursor.isEnd
}

func (t *Trie) prefix(prfx string) *Node {
	wLen := len(prfx)
	cursor := t.root
	for i := 0; i < wLen; i++ {
		// offsetting with 'a' based on ascii tables
		charIx := (strings.ToLower(prfx)[i]) - offsetASCIIa
		// empty char --> time to append
		if cursor.children[charIx] == nil {
			return nil
		}
		// update cursor for traversal
		cursor = cursor.children[charIx]
	}
	return cursor
}

// FindWords locates n words based on prfx
func (t *Trie) FindWords(prfx string, n int) []string {
	res := make([]string, 0, n)
	// find node where prefix ends
	prfxNode := t.prefix(prfx)
	if prfxNode == nil {
		// no prefix node means there is no words
		return nil
	}
	if prfxNode.isEnd {
		res = append(res, prfx)
	}

	var recurse func(node *Node, prfx string)
	recurse = func(node *Node, prfx string) {
		for i, c := range node.children {
			if c == nil {
				continue
			}
			if len(res) == n {
				return
			}

			prfx += string(i + offsetASCIIa)
			// word found
			if c.isEnd {
				res = append(res, prfx)
			}

			// backtracking-like logic
			if len(res) < n {
				recurse(c, prfx)
			}
			prfx = prfx[:len(prfx)-1]
		}
	}

	recurse(prfxNode, prfx)
	return res
}

// TODO func that returns a 2d slice with suggestions for each char from input:
// https://leetcode.com/problems/search-suggestions-system/description/?envType=study-plan-v2&envId=leetcode-75
func (t *Trie) GetSuggestions(word string, n int) [][]string {
	res := make([][]string, 0, len(word))
	for i := 1; i <= len(word); i++ {
		res = append(res, t.FindWords(word[:i], n))
	}
	return res
}
