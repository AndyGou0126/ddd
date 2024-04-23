package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type WordCountPair struct {
    Word  string
    Count int
}

func main() {
    // Read the file content
    fileContent, err := ioutil.ReadFile("900.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    // Split the content into words
    words := strings.Fields(string(fileContent))

    for k, word := range words {
        words[k] = word
    }
    // Create a map to store word counts
    wordCounts := make(map[string]int)
    for _, word := range words {
        for i:=0;i<3;i++ {
            word = strings.Trim(word, ",")
            word = strings.Trim(word, ".")
            word = strings.Trim(word, "(")
            word = strings.Trim(word, ")")
            word = strings.Trim(word, "\"")
            word = strings.Trim(word, "?")
            word = strings.Trim(word, "!")
            word = strings.Trim(word, ";")
            word = strings.Trim(word, ":")
            word = strings.Trim(word, "-")
        
            word = strings.Trim(word, "0")
            word = strings.Trim(word, "1")
            word = strings.Trim(word, "2")
            word = strings.Trim(word, "3")
            word = strings.Trim(word, "4")
            word = strings.Trim(word, "5")
            word = strings.Trim(word, "6")
            word = strings.Trim(word, "7")
            word = strings.Trim(word, "8")
            word = strings.Trim(word, "9")
        }
        
        wordCounts[word]++
    }

    // Convert word-count map to a slice of WordCountPair structs
    wordCountPairs := make([]WordCountPair, 0, len(wordCounts))
    for word, count := range wordCounts {
        wordCountPairs = append(wordCountPairs, WordCountPair{word, count})
    }

    // Sort the word-count pairs by count in descending order
    sort.Slice(wordCountPairs, func(i, j int) bool {
        return wordCountPairs[i].Count > wordCountPairs[j].Count
    })

    // Write the sorted word-count pairs to a new file
    outputFile, err := os.Create("sorted_word_counts.txt")
    if err != nil {
        fmt.Println("Error creating output file:", err)
        return
    }
    defer outputFile.Close()

    writer := bufio.NewWriter(outputFile)
    for _, pair := range wordCountPairs {
        _, err := writer.WriteString(fmt.Sprintf("%s\n", pair.Word))
        if err != nil {
            fmt.Println("Error writing to output file:", err)
            return
        }
    }
    writer.Flush()

    fmt.Println("Word counts sorted and written to sorted_word_counts.txt")
}