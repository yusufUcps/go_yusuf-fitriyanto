package main

import (
    "fmt"
    "strings"
    "sync"
)

func main() {
    text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
    var m sync.Mutex
    var wg sync.WaitGroup

    results := make(chan map[string]int)

    for _, word := range strings.Split(text, " ") {
        wg.Add(1)
        go func(word string) {
            frekuensi := make(map[string]int)
            for _, letter := range word {
                frekuensi[string(letter)]++
            }
            results <- frekuensi
            wg.Done()
        }(word)
    }

    go func() {
        wg.Wait()
        close(results)
    }()

    finalFrequency := make(map[string]int)

    for result := range results {
        m.Lock()
        for letter, count := range result {
            finalFrequency[letter] += count
        }
        m.Unlock()
    }

    for letter, count := range finalFrequency {
        fmt.Printf("%s: %d\n", letter, count)
    }
}
