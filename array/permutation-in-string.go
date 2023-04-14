package array

// 567 字符串的排列
func checkInclusion(s1 string,s2 string) bool  {
	window,need,left ,right,valid := make(map[byte]int),make(map[byte]int),0,0,0

	for i := range s1 {
		need[s1[i]]++
	}

	for right < len(s2) {
		if _,ok := need[s2[right]] ; ok {
			window[s2[right]]++
			if window[s2[right]] == need[s2[right]] {
				valid++
			}
		}
		right++

		if right - left == len(s1) {
			if valid == len(need){
				return true
			}

			if _,ok := need[s2[left]] ; ok {
				if window[s2[left]] == need[s2[left]] {
					valid--
				}
				window[s2[left]]--
			}
			left++
		}
	}
	return false
}
