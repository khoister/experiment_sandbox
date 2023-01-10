package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strings"
)

type AnagramFinder struct {
    table map[string][]string
}

// Creates an empty trie
func NewAnagramFinder() *AnagramFinder {
    return &AnagramFinder { table: make(map[string][]string)}
}

func getSignature(word string) string {
    key := ""
    if len(word) > 0 {
        arr := strings.Split(word, "")
        sort.Strings(arr)
        key = strings.Join(arr, "")
    }
    return key
}

// Insert a word to a set of anagrams
func (finder *AnagramFinder) add(word string) {
    key := getSignature(word)
    if len(key) > 0 {
        _, ok := finder.table[key]
        if !ok {
            finder.table[key] = []string{}
        }
        finder.table[key] = append(finder.table[key], word)
    }
}

// Find anagrams for a certain word
func (finder *AnagramFinder) find(word string) []string {
    key := getSignature(word)
    if len(key) > 0 {
        anagrams, ok := finder.table[key]
        if ok {
            return anagrams
        }
    }
    return []string{}
}

// Load words from a file into a trie data structure
func loadAnagrams(filename string) *AnagramFinder {
    readFile, err := os.Open(filename)

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)

    anagramFinder := NewAnagramFinder()
    for fileScanner.Scan() {
        word := fileScanner.Text()
        anagramFinder.add(word)
    }
    readFile.Close()
    return anagramFinder
}
