package main

import (
    "bufio"
    "fmt"
    "os"
)

// Trie data structure for finding common prefixes
type Trie struct {
    root *TrieNode
}

// Creates an empty trie
func NewTrie() *Trie {
    return &Trie {NewTrieNode(0)}
}

// Add a word to the trie
func (trie *Trie) add(word string) {
    node := trie.root
    for _, c := range []byte(word) {
        node = node.addChild(c)
    }
    node.isWord = true
}

// Find all words in the trie with the same prefix
func (trie *Trie) similar(prefix string) []string {
    return suffix(trie.get(prefix), prefix, []string{})
}

// Get the "lowest" node in the trie for a particular prefix string
func (trie *Trie) get(prefix string) *TrieNode {
    node := trie.root
    for _, c := range []byte(prefix) {
        childNode := node.getChild(c)
        if childNode == nil {
            return nil
        }
        node = childNode
    }
    return node
}

// Retrieves all the complete words in the trie that begins with the same prefix
func suffix(node *TrieNode, prefix string, similarWords []string) []string {
    if len(prefix) == 0 {
        return []string{}
    }
    if node != nil {
        if node.isWord {
            similarWords = append(similarWords, prefix)
        }
        for ch, childNode := range node.children {
            similarWords = suffix(childNode, prefix + string(ch), similarWords)
        }
    }
    return similarWords
}

// Load words from a file into a trie data structure
func loadTrie(filename string) *Trie {
    readFile, err := os.Open(filename)

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

    trie := NewTrie()
    for fileScanner.Scan() {
        word := fileScanner.Text()
        trie.add(word)
    }
    readFile.Close()
    return trie
}

