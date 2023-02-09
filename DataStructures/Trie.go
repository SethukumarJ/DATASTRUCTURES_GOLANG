package main

import "fmt"

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children map[rune]*TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

func (t *Trie) PopulateSuffixTrie(str string) {
	for i := 0; i < len(str); i++ {
		t.InsertSubStringStartingAt(i, str)
	}
}

func (t *Trie) InsertSubStringStartingAt(index int, str string) {
	node := t.root
	for i := index; i < len(str); i++ {
		letter := rune(str[i])
		child, exists := node.children[letter]
		if !exists {
			child = NewTrieNode()
			node.children[letter] = child
		}
		node = child
	}
}

func (t *Trie) Contains(str string) bool {
	node := t.root
	for _, letter := range str {
		child, exists := node.children[letter]
		if !exists {
			return false
		}
		node = child
	}
	return true
}

func (t *Trie) PrintWordsStartingWith(str string) {
	node := t.root
	for _, letter := range str {
		child, exists := node.children[letter]
		if !exists {
			return
		}
		node = child
	}
	t.dfs(node, str)
}

func (t *Trie) dfs(node *TrieNode, word string) {
	if len(node.children) == 0 {
		fmt.Println(word)
		return
	}
	for letter, child := range node.children {
		t.dfs(child, word+string(letter))
	}
}

func main() {
	trie := Trie{root: NewTrieNode()}
	trie.PopulateSuffixTrie("annan")
	trie.InsertSubStringStartingAt(0,"and")
	fmt.Println(trie.Contains("hello")) // false
	fmt.Println(trie.Contains("ann")) 
	trie.PrintWordsStartingWith("an")  // true
}

