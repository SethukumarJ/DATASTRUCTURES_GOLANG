package main

import "fmt"

type Trie struct {
	root *TrieNode
}
type TrieNode struct {
	children map[interface{}]*TrieNode
}

const endSymbol = "*"

func (t *Trie) populateSuffixTrie(str string) {

	for i := 0; i < len(str); i++ {
		t.insertSubStringStartingAt(i, str)
	}
}

func (t *Trie) insertSubStringStartingAt(index int, str string) {

	node := t.root
	for i := index; i < len(str); i++ {

		letter := str[i]
		_, exists := node.children[letter]
		if !exists {
			newNode := new(TrieNode)
			node.children[letter] = newNode
		}
		node = node.children[letter]

	}
	node.children[endSymbol] = nil
}

func (t *Trie) contains(str string) bool {

	node := t.root
	for i := 0; i < len(str); i++ {
		letter := str[i]
		_, exists := node.children[letter]
		if !exists {
			return false
		}
		node = node.children[letter]
	}

	_, exists := node.children[endSymbol]
	return exists
}

func main() {

	trie := Trie{}
	trie.populateSuffixTrie("annan")

	fmt.Println(trie.contains("hello"))
}
