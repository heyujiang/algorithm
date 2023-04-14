package array

import "strings"


//151 . 反转字符串中的单词
func reverseWords(s string) string {
	res := strings.Split(s," ")

	slow,fast := 0,0
	for fast < len(res) {
		if res[fast] != "" {
			res[slow] = res[fast]
			slow++
		}
		fast++
	}
	res = res[:slow]

	left ,right := 0,len(res)-1
	for left<right{
		res[left],res[right] = res[right],res[left]
		left++
		right--
	}
	return strings.Join(res," ")
}
