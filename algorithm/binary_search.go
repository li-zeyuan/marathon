package algorithm

/*
二分法通过对折的方式查找一个数据，条件是必须是一个有序的数组。数组底层是顺序链表，可以随机获取一个值 O (1)
参考
- https://learnku.com/articles/45836
*/

type BinarySearchObj struct {
	l      []int
	target int
}

func NewBSO(l []int, target int) *BinarySearchObj {
	return &BinarySearchObj{
		l:      l,
		target: target,
	}
}

func (b *BinarySearchObj) BinarySearch() int {
	l, r := 0, len(b.l)-1
	for l <= r {
		mid := (l + r) / 2
		if b.l[mid] == b.target {
			return 0
		}
		if b.target > b.l[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}
