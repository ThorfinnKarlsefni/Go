package leetcode

import "sort"

func FindMedianSortedArrays(nums1 []int, nums2 []int) (nums float64) {
	arr := append(nums1, nums2)
	sort.Ints(arr)
	if len(arr)%2 == 0 {
		index := len(arr) / 2
		nums = float64(arr[index]+arr[index-1]) / 2
	} else {
		index := len(arr) / 2
		nums = float64(arr[index])
	}

	return nums
}

func append(nums1, nums2 []int) []int {
	arr := make([]int, len(nums1)+len(nums2))
	at := copy(arr, nums1[:len(nums1)])
	copy(arr[at:], nums2[:len(nums2)])
	return arr
}
