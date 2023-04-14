package array

//28. 找出字符串中第一个匹配项的下标
func strStr(haystack string, needle string) int {
	left,right := 0,0
	for right < len(haystack) {
		right++

		if right - left == len(needle) {
			if haystack[left:right] == needle {
				return left
			}
			left++
		}
	}
	return -1
}
