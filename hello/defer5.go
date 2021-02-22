package main

import (
	"errors"
	"fmt"
	"math"
	"unsafe"
)

func divide(x, y int) (int, error) {
	if y <= 0 {
		return 0, errors.New("cant divide with 0")
	}
	return x / y, nil
}

func setZeropassbyval(i int) {
	i = 0
	fmt.Println("Set Zero pass by value: ", i)
}

func setZeropassbyref(i *int) {
	*i = 0
	fmt.Println("Set zero pass by references: ", *i)
}

func init() {
	fmt.Println("init ...")
}

type developer struct {
	skill    []string
	exp      int
	salary   int
	location string
	line     string
	company  string
}

// type rect struct {
// 	width, height int
// }

//  func (r *rect) area() int {
//  	return r.width * r.height
//  }

//  func (r rect) perim() int {
//  	return 2*r.width + 2*r.height
//  }

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("area : ", g.area())
	fmt.Println("perim : ", g.perim())
}

func main() {
	// defer fmt.Println("after end of main func")
	// defer fmt.Println("time for bye!")
	// fmt.Println("hello start")
	fmt.Println("--------------")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered : syatem cant find")
			fmt.Println("----------------")
		}
	}()

	// path := "user/hack/paa/err.txt"
	// _, err := os.Stat(path)
	// if os.IsNotExist(err) {
	// 	file, err := os.Create(path)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	defer file.Close()
	// }
	//fmt.Println("done creating file", path)
	i := 5
	fmt.Println("value i ", i)
	fmt.Println("address value i ", &i)
	p := &i
	fmt.Println("value p ", p)
	fmt.Println("address value p ", &p)
	fmt.Println("value *p ", *p)
	fmt.Println("size p ", unsafe.Sizeof(p))
	fmt.Println("-------------")
	i = 5
	fmt.Println("value i ", i)
	setZeropassbyval(i)
	fmt.Println("value i: ", i)
	fmt.Println("--------------")
	fmt.Println("value of i: ", i)
	fmt.Println("value of &i", &i)
	setZeropassbyref(&i)
	fmt.Println("value if i:", i)
	fmt.Println("--------------")
	fmt.Println(developer{})
	dev := developer{
		skill:    []string{"node.js", "angula"},
		exp:      0,
		location: "utchajs",
		line:     "na",
		company:  "delta",
	}
	dev.salary = 9000
	fmt.Println(dev)
	fmt.Println("--------------")
	devp := &dev
	fmt.Println(devp)
	fmt.Println(dev.company)
	fmt.Println(devp.company)
	devp.company = "alt f4"
	fmt.Println(devp.company)
	dev.company = "aaa"
	fmt.Println(dev)
	fmt.Println("--------------")
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())
	fmt.Println("--------------")
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
	fmt.Println("--------------")
	r = rect{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)
	fmt.Println("--------------")
	if rr, ee := divide(10, 0); ee != nil {
		fmt.Println("error :", ee)
	} else {
		fmt.Println("result :", rr)
	}
}
