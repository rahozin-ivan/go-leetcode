package leetcode

func lengthOfLongestSubstring(s string) int {
	var max, start int
	seen := make(map[rune]int)
	for i, r := range s {
		if seen[r] > start {
			start = seen[r]
		}
		if i-start+1 > max {
			max = i - start + 1
		}
		seen[r] = i + 1
	}
	return max
}
