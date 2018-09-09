package leetcode

import (
	"testing"
)

func TestReplaceWords(t *testing.T) {
	dict := []string{"cat", "bat", "rat"}
	sentence := "the cattle was rattled by the battery"
	output := "the cat was rat by the bat"
	if r := replaceWords(dict, sentence); r != output {
		t.Errorf("expected %v, got %v", output, r)
	}
}
