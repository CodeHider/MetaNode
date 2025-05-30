package chapter4

import "sort"

// RemoveDuplicates
/**
删除数组中的重复项.返回数组中不重复元素个数
*/
func RemoveDuplicates(nums []int) int {
	//从小到大排序
	sort.Ints(nums)
	i := 1
	// i-1是最后一个不重复的数据
	for j := 0; j < len(nums); j++ {
		if nums[j] != nums[i-1] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
