package leetcode

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strs {
		// 1. sort.Slice could be used
		// 2. sum could be generated
		k := sortString(s)
		if _, ok := m[k]; !ok {
			m[k] = make([]string, 0)
		}

		m[k] = append(m[k], s)
	}

	var result [][]string
	for _, v := range m {
		result = append(result, v)
	}

	return result
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
