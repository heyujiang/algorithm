package array

//3. 无重复字符的最长子串

func lengthOfLongestSubstring(s string) int {
	window,left,right,res :=make(map[byte]int), 0,0,0

	for right < len(s) {
		c := s[right]
		right++
		window[c]++
		for window[c] > 1 {
			window[s[left]]--
			left++
		}

		if right - left > res {
			res = right - left
		}
	}
	return res
}
