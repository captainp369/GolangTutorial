package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init ...")
}

func main() {

	fmt.Println("Start")

	var vartypebool bool = true
	var vartypeint int = 3
	vartypefloat := 11.42
	const xy int = 11

	fmt.Println(vartypebool)
	fmt.Println(vartypeint)
	fmt.Println(vartypefloat)
	fmt.Println(xy)

	for i := 0; i <= 20; i++ {
		if i == 5 {
			continue
		} else if i == 12 {
			break
		}
		fmt.Println(i)
	}

	sum := 0
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum)

}
