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
