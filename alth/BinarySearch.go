package alth

/**
如题二分查找法go语言实现
*/
type BinarySearch struct{}

func (BinarySearch BinarySearch) Search(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	return bs(arr, target, 0, n-1)

}

func bs(arr []int, target int, start int, end int) int {

	if start > end {
		return -1
	}
	middle := (start + end) / 2
	if arr[middle] == target {
		return middle
	}
	if arr[middle] > target {
		return bs(arr, target, start, middle-1)
	} else {
		return bs(arr, target, middle+1, end)
	}
}
