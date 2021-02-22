package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func prism(lenght, width, height int) int {
	return lenght * width * height
}

func findFirstandlast(msg string) (string, string) {
	return string(msg[0]), string(msg[len(msg)-1])
}

func swap(x, y int) (a, b int) {
	a = y
	b = x
	return
}

func sum(nums ...int) {
	fmt.Print(nums, "=")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func factusingfor(n int) int {
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	return factorial
}

func init() {
	fmt.Println("init start .......")
}

func main() {
	aa, bb := 11, 67
	res := plus(aa, bb)
	fmt.Println(aa, "+", bb, "=", plus(aa, bb))
	res = prism(12, 42, 111)
	fmt.Println("prism volume =", res)
	fmt.Println("---------")
	msng := "hello world whis is hajsd a"
	a, b := findFirstandlast(msng)
	fmt.Println(msng, "=>", a, ",", b)
	i, j := 1, 2
	i, j = swap(i, j)
	fmt.Println(i, j)
	fmt.Println("---------")
	nums := []int{1, 2, 3, 4, 5, 6}
	sum(1, 2)
	sum(1, 99, 100)
	sum(nums...)
	sum(nums[0:4]...)
	fmt.Println("---------")
	nextint := intSeq()
	fmt.Println(nextint())
	fmt.Println(nextint())
	fmt.Println(nextint())
	fmt.Println(nextint())
	nextints := intSeq()
	fmt.Println(nextints())
	fmt.Println(nextints())
	fmt.Println(nextint())
	fmt.Println("---------")
	fmt.Println(fact(10))
	fmt.Println(factusingfor(10))
}
