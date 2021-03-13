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

时间复杂度：O(nlogn)

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

// =============================
/*
冒泡排序
时间复杂度：O(n2)

思路
1、依次比较相邻的俩个数，交换大的数到后面，一轮后，最大的数在最右边
2、重复1的，遍历比较除最右边元素外的其他数
*/
func BubbleSort(ls []int) {
	l := len(ls)
	if l < 1 {
		return
	}

	for i := 0; i < l; i++ {
		for j := 0; j < l-i-1; j++ { // 遍历比较除最右边已排序元素外的其他数
			if ls[j+1] < ls[j] {
				ls[j], ls[j+1] = ls[j+1], ls[j] // 交换大的元素到右侧
			}
		}
	}
}

// =====================
/*
选择排序：剩下的元素，假定第一个为最小，遍历，找到最小，与剩下元素的第一个交换
时间复杂度：O(n2)
思路
1、假定剩下元素的第一个未最小
2、依次遍历剩下元素，若有比它小的，则交换位置
*/
func SelectSort(ls []int) {
	l := len(ls)
	if l < 1 {
		return
	}

	for i := 0; i < len(ls); i++ {
		for j := i + 1; j < l; j++ { // 遍历右侧剩下的元素
			if ls[i] > ls[j] {
				ls[j], ls[i] = ls[i], ls[j] // 交换最小的元素到第一个位置
			}
		}
	}
}
