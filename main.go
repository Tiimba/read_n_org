package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	userPath := os.Args[1]
	file, err := os.Open(userPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	wordCount := make(map[string]int)
	re := regexp.MustCompile("[^a-zA-Z ]+")

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		line = re.ReplaceAllString(line, "")

		words := strings.Split(line, " ")
		for _, word := range words {
			if word != "" {
				wordCount[word]++
			}
		}
	}
	pairs := make([][2]interface{}, 0, len(wordCount))
	for word, count := range wordCount {
		pairs = append(pairs, [2]interface{}{word, count})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0].(string) < pairs[j][0].(string)
	})
	for _, pair := range pairs {
		fmt.Printf("%s: %d\n", pair[0].(string), pair[1].(int))
	}
}
