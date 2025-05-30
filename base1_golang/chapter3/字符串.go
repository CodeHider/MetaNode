package chapter3

/*
*
**  栈的使用

	有效的括号

题目：给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
*/
func IsValid(s string) bool {
	// ((((
	var staticMap = map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	//初始化切片
	var charArr = make([]rune, 0)
	for _, value := range s {
		if '(' == value || '{' == value || '[' == value {
			charArr = append(charArr, value)
		} else {
			if len(charArr) > 0 {
				if staticMap[value] != charArr[len(charArr)-1] {
					return false
				} else {
					charArr = charArr[0 : len(charArr)-1]
				}
			} else {
				return false
			}
		}
	}
	if len(charArr) == 0 {
		return true
	}
	return false
}

/*
*
最长公共前缀

编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。
*/
func LongestCommonPrefix(strs []string) string {
	//k:表示字符的下标位置,出现了几次. index_c
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		prefix = getSingleLongest(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func getSingleLongest(prefix string, str2 string) string {
	length := min(len(prefix), len(str2))
	index := 0
	for index < length && prefix[index] == str2[index] {
		index++
	}
	//在之前已经截取的字符串中的前index字符
	return prefix[:index]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
