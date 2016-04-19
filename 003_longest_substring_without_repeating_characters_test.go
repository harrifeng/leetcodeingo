package leetcodeingo

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	if lengthOfLongestSubstring("abcabcbb") != 3 {
		t.Error("error occur")
	}

	if lengthOfLongestSubstring("bbbb") != 1 {
		t.Error("error occur")
	}

	if lengthOfLongestSubstring("abab") != 2 {
		t.Error("error occur")
	}

}
