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

/*
*

	加一
*/
func PlusOne(digits []int) []int {
	//[0，0，0]
	last := len(digits) - 1
	mod := 1
	over := false
	for i := last; i >= 0; i-- {
		temp := digits[i] + mod
		if temp == 10 {
			digits[i] = 0
			mod = 1
		} else {
			digits[i] = temp
			mod = 0
		}
		if i == 0 && mod == 1 {
			over = true
		}
	}
	if over {
		length := len(digits) + 1
		digits = make([]int, length)
		digits[0] = 1
	}
	return digits
}

/*
*

	一个整数目标值 target
*/
func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
