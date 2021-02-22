package main

import (
	"fmt"
)

func init() {
	fmt.Println("start")
	if i := 0; i < 5 {
		fmt.Println("-----------")
	}
}

func main() {
	fmt.Println("daylearn")

	i := 3
	fmt.Println("write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
		fmt.Println("==========")
		k := 1
		switch k {
		case 2:
			fmt.Println("two")
			fallthrough
		case 1:
			fmt.Println("one")
			fallthrough
		case 3:
			fmt.Println("three")
			fallthrough
		default:
			fmt.Println("default...")
		}
	}

	var a [3]int
	b := [5]int{1, 2, 3, 4, 5}
	a[0] = 1
	a[1] = 2
	a[2] = 43
	fmt.Println("lenght of b = ", len(b))
	fmt.Println("array a =", a)
	var threed [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			threed[i][j] = i + j
		}
	}
	fmt.Println(threed)
	fmt.Println("------------")

	aa := make([]string, 3)
	aa[0] = "a"
	aa[1] = "b"
	aa[2] = "c"
	fmt.Print(aa)
	fmt.Println("---")
	ab := make([]string, len(aa))
	copy(ab, aa)
	ab = append(aa, "a", "b")
	fmt.Println(ab)
	ac := ab[0 : len(ab)-1]
	fmt.Println("value of slice ac:", ac)
	ac[0] = "x"
	fmt.Println("new value of slice ac = ", ac)
	fmt.Println("--------------")
	m := make(map[string]int)
	m["math"] = 89
	m["sci"] = 99
	m["tees"] = 0
	fmt.Println("lenght of map", len(m))
	fmt.Println(m)
	fmt.Println("-------------")
	for i, num := range m {
		fmt.Println(i, num)
	}
	for i, num := range b {
		fmt.Println(i, num)
	}
}
