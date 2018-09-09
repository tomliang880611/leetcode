package leetcode

import (
	"sort"
	"strings"
)

func replaceWords(dict []string, sentence string) string {
	// sort the dict by len in asecent order
	sort.Slice(dict, func(i, j int) bool {
		return len(dict[i]) > len(dict[j])
	})

	words := strings.Split(sentence, " ")
	for i, word := range words {
	next:
		for _, d := range dict {
			// replace and skip match for longer dict/prefix
			if strings.HasPrefix(word, d) {
				words[i] = d
				continue next
			}
		}
	}

	return strings.Join(words, " ")
}
