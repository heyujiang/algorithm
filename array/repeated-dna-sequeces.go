package array

// 187 . 重复的DNA序列
func findRepeatedDNASequences(s string) []string {
	strNums,left,right,res := make(map[string]int),0,0,make([]string,0)

	for right < len(s) {
		right++
		for right - left == 10  {
			strNums[s[left:right]]++
			left++
		}
	}

	for str,num := range strNums{
		if num > 1 {
			res = append(res,str)
		}
	}
	return res
}