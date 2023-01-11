package main

import (
    "database/sql"
    "fmt"
    "strings"
    _ "modernc.org/sqlite"
)

type Dictionary struct {
    db *sql.DB
}

func loadDictionary(dbfile string) *Dictionary{
    db, err := sql.Open("sqlite", dbfile)
    checkError(err)
    return &Dictionary{db}
}

func (dictionary *Dictionary) find(word string) []string {
    queryString := fmt.Sprintf("SELECT part_of_speech, definition FROM dictionary WHERE word = '%s' ORDER BY part_of_speech", strings.ToLower(word))
    rows, err := dictionary.db.Query(queryString)
    checkError(err)

    var partOfSpeech string
    var definition string

    definitions := []string{}
    for rows.Next() {
        err = rows.Scan(&partOfSpeech, &definition)
        checkError(err)

        row := fmt.Sprintf("%s: %s", partOfSpeech, definition)
        definitions = append(definitions, row)
    }
    rows.Close()
    return definitions
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

