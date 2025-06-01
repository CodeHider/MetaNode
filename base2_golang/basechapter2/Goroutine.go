package basechapter2

import "fmt"

/*
*

	      编写一个程序，
		  使用 go 关键字启动两个协程，
	      一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
	      go 关键字的使用、协程的并发执行。
*/
func GoRutineFunc() {
	go func() {
		size := 10
		for i := 0; i < size; i++ {
			if i%2 == 0 {
				fmt.Println("打印偶数", i)
			}
		}
	}()

	go func() {
		size := 10
		for i := 0; i < size; i++ {
			if i%2 != 0 {
				fmt.Println("打印寄数", i)
			}
		}
	}()
}
