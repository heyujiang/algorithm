package array

//438. 找到字符串中所有字母异位词
func findAnagrams(s string,p string) []int {
	windows,need,left,right,valid,res := make(map[byte]int),make(map[byte]int),0,0,0,make([]int,0)

	for i := range p {
		need[p[i]]++
	}

	for right < len(s) {
		if _,ok := need[s[right]];ok {
			windows[s[right]]++
			if windows[s[right]] == need[s[right]] {
				valid++
			}
		}
		right++

		if right - left == len(p) {
			if valid == len(need) {
				res = append(res,left)
			}

			if _,ok := need[s[left]];ok {
				if windows[s[left]] == need[s[left]]{
					valid--
				}
				windows[s[left]]--
			}

			left++
		}
	}

	return res

}
