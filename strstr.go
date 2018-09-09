package leetcode

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	nl, hl, cur := len(needle), len(haystack), 0

	for ; cur < hl; cur++ {
		i := 0
		// use slicing in golang could speed this up
		for ; i < nl && cur+i < hl && haystack[cur+i] == needle[i]; i++ {
		}
		if i == nl {
			return cur
		}
	}

	return -1
}
