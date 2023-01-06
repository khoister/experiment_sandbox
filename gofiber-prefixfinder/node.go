package main

type TrieNode struct {
    char byte
    isWord bool
    children map[byte]*TrieNode
}

func NewTrieNode(ch byte) *TrieNode {
    return &TrieNode{char: ch, isWord: false, children: make(map[byte]*TrieNode)}
}

func (node *TrieNode) addChild(ch byte) *TrieNode {
    if !node.hasChild(ch) {
        node.children[ch] = NewTrieNode(ch)
    }
    return node.getChild(ch)
}

func (node *TrieNode) hasChild(ch byte) bool {
    return node.getChild(ch) != nil
}

func (node *TrieNode) getChild(ch byte) *TrieNode {
    child, ok := node.children[ch]
    if !ok {
        return nil
    }
    return child
}
