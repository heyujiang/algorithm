package array

//5 .  最长回文子串

func longestPalindromic(s string) string  {
	res := ""
	for i:=0;i<len(s);i++ {
		s1 := palindromic(s,i,i)
		s2 := palindromic(s,i,i+1)

		if len(res) < len(s1){
			res = s1
		}
		if len(res) < len(s2){
			res = s2
		}
	}
	return res
}

func palindromic(s string, left ,right int ) string  {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return s[left+1:right]
}
