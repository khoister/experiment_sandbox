package main

import (
    "strings"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/pug"
)

func main() {
    anagramFinder := loadAnagrams("enable1.txt")
    trie := loadTrie("enable1.txt")
    engine := pug.New("./views", ".pug")
    app := fiber.New(fiber.Config{
        Views: engine,
    })

    app.Get("/", func(c *fiber.Ctx) error {
        return c.Render("index", fiber.Map{"Title": "Word Tools"})
    })

    app.Get("/prefix", func(c *fiber.Ctx) error {
        similarWords := trie.similar(strings.ToLower(c.Query("search")))
        Jade_similarwords(similarWords, c)
        return nil
    })

    app.Get("/anagram", func(c *fiber.Ctx) error {
        anagrams := anagramFinder.find(strings.ToLower(c.Query("search")))
        Jade_similarwords(anagrams, c)
        return nil
    })

    app.Listen(":3000")
}
