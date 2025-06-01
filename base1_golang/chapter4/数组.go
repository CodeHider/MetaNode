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

// Merge /*
/**
合并区间：
	以数组 intervals 表示若干个区间的集合，
	其中单个区间为 intervals[i] = [starti, endi] 。
	请你合并所有重叠的区间，
	并返回 一个不重叠的区间数组，
	该数组需恰好覆盖输入中的所有区间 。
*/
func Merge(intervals [][]int) [][]int {
	//输入: intervals = [[1,3],[2,6],[8,10],[15,18]]
	//输出: [[1,6],[8,10],[15,18]]
	temp := make([]int, 0)
	temp = intervals[0]
	sort.Ints(temp)
	if len(intervals) <= 1 {
		return intervals
	}
	//切片的长度
	result := make([][]int, 0)
	result = append(result, temp)
	for i := 1; i < len(intervals); i++ {
		arr := intervals[i]
		sort.Ints(arr)
		//相交
		if temp[0] >= arr[0] && temp[0] <= arr[len(arr)-1] {
			temp[0] = arr[0]
			if temp[len(temp)-1] < arr[len(arr)-1] {
				temp[len(temp)-1] = arr[len(arr)-1]
			}
		} else if temp[0] <= arr[0] && arr[0] <= temp[len(temp)-1] {
			//相交
			if temp[len(temp)-1] < arr[len(arr)-1] {
				temp[len(temp)-1] = arr[len(arr)-1]
			}
		} else {
			//不相交
			result = append(result, arr)
		}
	}
	return result
}
