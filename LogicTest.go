package main

import (
	"fmt"
	"strings"
)

func isAnagram(arr []string) [][]string {
	anagramGroups := make(map[string][]string)
	for _, word := range arr {
		sorted := sortString(word)
		anagramGroups[sorted] = append(anagramGroups[sorted], word)
	}

	var result [][]string
	for _, group := range anagramGroups {
		result = append(result, group)
	}

	return result
}

func sortString(s string) string {
	s = strings.ToLower(s)
	r := []rune(s)
	for i := 0; i < len(r)-1; i++ {
		for j := i + 1; j < len(r); j++ {
			if r[i] > r[j] {
				r[i], r[j] = r[j], r[i]
			}
		}
	}
	return string(r)
}

func main() {
	arr := []string{"cook", "save", "taste", "aves", "vase", "state", "map"}
	fmt.Println(isAnagram(arr))
}