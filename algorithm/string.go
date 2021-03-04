package algorithm

import "strings"

/*
请实现一个算法，确定一个字符串的所有字符【是否全都不同】。这里我们要求【不允许使用额外的存储结构】。
给定一个string，请返回一个bool值,true代表所有字符全都不同，false代表存在相同的字符。
保证字符串中的字符为【ASCII字符】。字符串的长度小于等于【3000】。

思路
	- strings.Count的使用
 */
func isUniqueString(s string) bool {
	if strings.Count(s,"") > 3000{
		return  false
	}
	for _,v := range s {
		if v > 127 {
			return false
		}
		if strings.Count(s,string(v)) > 1 {
			return false
		}
	}
	return true
}

/*
请实现一个算法，在不使用【额外数据结构和储存空间】的情况下，翻转一个给定的字符串(可以使用单个过程变量)。
给定一个string，请返回一个string，为翻转后的字符串。保证字符串的长度小于等于5000。

思路
	- 将string转成[]rune
	- 以中心字符为轴，交换两边的字符
 */

func ReversedString(s string) string {
	if len(s) > 5000 {
		return ""
	}

	strList := []rune(s)
	for i := 0; i < len(s)/2; i ++ {
		strList[i], strList[len(s)-1 - i] = strList[len(s)-1 - i], strList[i]
	}

	return string(strList)
}