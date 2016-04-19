package leetcodeingo

func lengthOfLongestSubstring(s string) int {
	last := 0
	maxv := 0
	m := make(map[rune]int)

	for i, c := range s {
		v, ok := m[c]
		if ok {
			if last < v+1 {
				last = v + 1
			}
		}
		if maxv < (i - last + 1) {
			maxv = i - last + 1
		}

		m[c] = i
	}
	return maxv
}
