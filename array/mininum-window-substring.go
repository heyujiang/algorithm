package array

// 76 最小覆盖子串
func minWindow(s string,t string) string  {
	window,need,res,left,right,valid := make(map[byte]int),make(map[byte]int),"",0,0,0
	for i:=0;i<len(t);i++ {
		need[t[i]]++
	}

	for right < len(s) {
		if _,ok := need[s[right]];ok {
			window[s[right]]++
			if window[s[right]] == need[s[right]] {
				valid++
			}
		}
		right++

		for left < right && valid == len(need) {
			if res == "" || len(res) > len(s[left:right]) {
				res = s[left:right]
			}

			if _,ok := need[s[left]];ok {
				if window[s[left]] == need[s[left]] {
					valid--
				}
				window[s[left]]--
			}
			left++
		}
	}

	return res
}