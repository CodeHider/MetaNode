package chapter1

import "sort"

/*
*
** 1.控制流程:
 136. 只出现一次的数字：
    给定一个非空整数数组，除了某个元素只出现一次以外，
    其余每个元素均出现两次。
    找出那个只出现了一次的元素。
    可以使用 for 循环遍历数组，
    结合 if 条件判断和 map 数据结构来解决，
    例如通过 map 记录每个元素出现的次数，
    然后再遍历 map 找到出现次数为1的元素。
*/
func singleNumber(nums []int) int {
	//创建一个空map
	numMap := make(map[int]int)
	for _, value := range nums {
		if count, ok := numMap[value]; ok {
			count += 1
			numMap[value] = count
		} else {
			numMap[value] = 1
		}
	}

	for key, value := range numMap {
		if value == 1 {
			return key
		}
	}
	//未找到数组个数为1个的元素
	return -1
}

/*
*
*
*  双指针法: 快慢指针
 */
func singleNumber1(nums []int) int {
	sort.Ints(nums)
	//慢指针,用指向第一个不重复的数字
	i := 0
	//统计数字出现的数量
	count := 0
	for k := 0; k < len(nums); k++ {
		if nums[i] != nums[k] {
			if count == 1 {
				return nums[i]
			}
			count = 1
			i = k
		} else {
			count++
		}
	}
	if count == 1 {
		return nums[i]
	}
	return -1
}
