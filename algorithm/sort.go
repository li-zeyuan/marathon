package algorithm

import "fmt"

type SortObj struct {
	List []int
}

/*
快排
1、选择列表的第一个元素为基准
2、头尾指针移动遍历列表，比基准小的在左边，比基准大的在右边
3、递归，重复1、2，排序基准左边、右边的子了列表

参考
https://learnku.com/articles/45802
*/
func (s *SortObj) QuicklySort() {

	fmt.Println(s.List)
}
