package algorithm

type SortObj struct {
	List []int
}

/*
快排
1、选择列表的第一个元素为基准
2、头尾指针移动遍历列表
3、尾指针元素 比 基准小，则对换位置，前移尾指针
4、头指针 比 头指针的后一个元素大，则对换位置，后移头指针
5、递归，重复1、2，排序基准左边、右边的子了列表

时间复杂度：

参考
https://learnku.com/articles/45802
*/
func (s *SortObj) QuicklySort() {
	if len(s.List) < 2 {
		return
	}

	l, r := 0, len(s.List)-1
	bValue := s.List[l] // 基准
	for l < r {
		if s.List[r] < bValue { // 尾指针元素 比 基准小，则对换位置，前移尾指针
			s.List[l], s.List[r] = s.List[r], s.List[l]
			r--
		} else if s.List[l] > s.List[l+1] { // 头指针 比 头指针的后一个元素大，则对换位置，后移头指针
			s.List[l], s.List[l+1] = s.List[l+1], s.List[l]
			l++
		} else {
			l++
		}
	}

	subLList := new(SortObj)
	subLList.List = s.List[:l]
	subLList.QuicklySort()

	subLRList := new(SortObj)
	subLRList.List = s.List[l+1:]
	subLRList.QuicklySort()
}
