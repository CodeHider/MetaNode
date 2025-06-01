package main

import (
	"MetaNode/base1_golang/chapter1"
	"MetaNode/base1_golang/chapter2"
	"MetaNode/base1_golang/chapter3"
	"MetaNode/base1_golang/chapter4"
	"MetaNode/base2_golang/basechapter1"
	"MetaNode/base2_golang/basechapter2"
	"fmt"
	"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	//1.只出现一次的数字：
	nums := []int{4, 1, 2, 1, 2}

	number := chapter1.SingleNumber(nums)

	fmt.Println(number)

	number1 := chapter1.SingleNumber1(nums)

	fmt.Println(number1)

	//2.回文数
	num := 12321
	palindrome := chapter2.IsPalindrome(num)
	fmt.Println("回文数", palindrome)

	//s = "(){}}{"
	s := "(){}}{"
	valid := chapter3.IsValid(s)
	fmt.Println("有效的括号", valid)

	strs := []string{"flower", "flow", "flight"}
	prefix := chapter3.LongestCommonPrefix(strs)
	fmt.Println("最长公共前缀", prefix)

	//longestPrefix := chapter3.LongestCommonPrefix_z(strs)
	//fmt.Println("最长公共前缀", longestPrefix)

	nums1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	duplicates := chapter4.RemoveDuplicates(nums1)
	fmt.Println("数组中的元素个数", duplicates)

	addOne := []int{9, 9, 9}
	one := chapter4.PlusOne(addOne)
	fmt.Println("加一的结果:", one)

	digits := []int{9}
	chapter4.PlusOne(digits)
	//nums = [2,7,11,15], target = 9
	twoSum := []int{2, 7, 11, 15}
	sum := chapter4.TwoSum(twoSum, 9)
	fmt.Println("数组的下标:", sum)

	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	chapter4.Merge(intervals)
	fmt.Println("合并之后的数组:", intervals)

	t := 10
	basechapter1.Increase(&t)
	fmt.Println("指针:", t)

	slice := []int{1, 2, 3, 4, 5}
	basechapter1.Multiply(&slice)
	fmt.Println("切片:", slice)

	basechapter2.GoRutineFunc()

	go

	//如何主函数的goRutine执行完毕了，则其他的goRutine也不会再执行了。
	time.Sleep(50 * time.Millisecond)

	//指针
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
}
