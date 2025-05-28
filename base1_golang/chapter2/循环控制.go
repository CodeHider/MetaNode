package chapter2

import "strconv"

// IsPalindrome /*
/**
回文数
  122221
*/
func IsPalindrome(x int) bool {
	str := strconv.Itoa(x)
	charArray := []byte(str)
	i := 0
	j := len(charArray) - 1
	for i < j {
		if charArray[i] != charArray[j] {
			return false
		} else {
			i++
			j--
		}
	}
	return true
}
