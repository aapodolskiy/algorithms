package main

import (
	"fmt"
)

////////////////////////////////
// implementation of MyTrie

type MyTrieInterface interface {
	Insert(word string)
	Search(word string) bool
	StartsWith(prefix string) bool
}

type MyTrieNode struct {
	children [26]*MyTrieNode
	isFinal  bool
}

type MyTrie struct {
	root *MyTrieNode
}

func MyTrieConstructor() MyTrieInterface {
	root := &MyTrieNode{[26]*MyTrieNode{}, false}
	return &MyTrie{root}
}

func _insertRecursive(node *MyTrieNode, word string) {
	index := int(word[0] - 'a')
	if node.children[index] == nil {
		node.children[index] = &MyTrieNode{[26]*MyTrieNode{}, false}
	}
	if len(word) == 1 {
		node.children[index].isFinal = true
		return
	}
	_insertRecursive(node.children[index], word[1:])
}

func (this *MyTrie) Insert(word string) {
	if len(word) > 0 {
		_insertRecursive(this.root, word)
	}
}

func _searchRecursive(node *MyTrieNode, word string) bool {
	index := int(word[0] - 'a')
	if node.children[index] == nil {
		return false
	}
	if len(word) == 1 {
		return node.children[index].isFinal
	}
	return _searchRecursive(node.children[index], word[1:])
}

func (this *MyTrie) Search(word string) bool {
	return _searchRecursive(this.root, word)
}

func _prefixSearchRecursive(node *MyTrieNode, prefix string) bool {
	index := int(prefix[0] - 'a')
	if node.children[index] == nil {
		return false
	}
	if len(prefix) == 1 {
		return true
	}
	return _prefixSearchRecursive(node.children[index], prefix[1:])
}

func (this *MyTrie) StartsWith(prefix string) bool {
	return _prefixSearchRecursive(this.root, prefix)
}

// end of MyTrie implementation
////////////////////////////////

func main() {
	trie := MyTrieConstructor()
	trie.Insert("a")
	trie.Insert("abctest")
	trie.Insert("xyzdskmk")
	trie.Insert("abcklm")

	fmt.Println("word", "\t", "trie.Search(word)", "\t", "trie.StartsWith(word)")
	for _, word := range []string{"a", "aa", "ab", "abc", "x", "xyz", "xyzdskmk"} {
		fmt.Println(word, "\t", trie.Search(word), "\t\t", trie.StartsWith(word))
	}
}
