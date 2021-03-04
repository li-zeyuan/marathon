package algorithm

/*
反转字符列表
使得{"A", "B", "C", "D"} -> {"D", "C", "B", "A"}
*/
func ReversedSlice(strList []string) []string {
	if len(strList) == 0 {
		return []string{}
	}

	for i := 0; i < len(strList)/2; i++ {
		strList[i], strList[len(strList)-1-i] = strList[len(strList)-1-i], strList[i]
	}

	return strList
}
