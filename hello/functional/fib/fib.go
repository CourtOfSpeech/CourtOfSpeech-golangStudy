package fib

//斐波那契数列
//1, 1, 2, 3, 5, 8, 13, 21, 34, 55
//   a, b
//      a, b
//func fibonacci() func() int {
func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
