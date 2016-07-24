package main

import (
	"fmt"
	"math"
	"time"
)

import "../rect"

type Person struct {
	name  string
	age   int
	email string
}

func main() {
	//变量和常量
	varDemo()
	//时间相关
	timeDemo()
	//数组例子
	arrayDemo()
	//map demo
	mapDemo()
	//for循环例子
	forDemo()
	//new和make使用和区别
	newMakeDemo()
	//函数测试，有一个返回值
	fmt.Println("多数求和函数", sum(12, 11, 4, 70))
	//返回多个结果的函数demo
	mutiReturnFunc()
	//函数闭包
	nextNumFunc := nextNum()
	for i := 0; i < 15; i++ {
		fmt.Println(nextNumFunc())
	}
	//结构体例子
	structDemo()

	shape.Interface_test()
	fmt.Println(math.Pi)

}

func structDemo() {
	fmt.Println("------------structDemo start------------")

	person := Person{"Tom", 30, "tom@gmail.com"}
	person = Person{name: "Tom", age: 30, email: "tom@gmail.com"}
	fmt.Println(person) //输出 {Tom 30 tom@gmail.com}

	pPerson := &person
	fmt.Println(pPerson) //输出 &{Tom 30 tom@gmail.com}

	pPerson.age = 40
	person.name = "Jerry"
	fmt.Println(person) //输出 {Jerry 40 tom@gmail.com}
	fmt.Println("------------structDemo end------------")

}

func nextNum() func() int {
	i, j := 1, 1
	return func() int {
		var tmp = i + j
		i, j = j, tmp
		return tmp
	}
}

func mutiReturnFunc() {
	fmt.Println("------------mutiReturn start------------")

	v, e := multiReturn("two")
	fmt.Println("返回正确结果:", v, e)
	v, e = multiReturn("four")
	fmt.Println("返回错误结果:", v, e)
	//通常的用法(注意分号后有e)
	if v, e = multiReturn("four"); e {
		// 正常返回
		fmt.Println("返回了错误结果:", v, e)
	} else {
		// 出错返回
		fmt.Println("返回了错误结果:", v, e)
	}
	fmt.Println("------------mutiReturn end------------")
}

/**
*new和make的使用及区别
 */
func newMakeDemo() {
	fmt.Println("------------new&make start------------")
	var v []int = make([]int, 10) // 切片v现在是对一个新的有10个整数的数组的引用
	fmt.Println("v []int = make([]int, 10):", v)

	// 不必要地使问题复杂化：
	var p *[]int = new([]int)
	fmt.Println(p) //输出：&[]

	*p = make([]int, 10, 100)
	fmt.Println(p)       //输出：&[0 0 0 0 0 0 0 0 0 0]
	fmt.Println((*p)[2]) //输出： 0

	// 习惯用法:
	v = make([]int, 10)
	fmt.Println(v) //输出：[0 0 0 0 0 0 0 0 0 0]
	fmt.Println("\n------------new&make end------------")

}
func mapDemo() {
	fmt.Println("\n------------map start------------")
	m := make(map[string]int)
	m["test1"] = 3
	m["test2"] = 18
	fmt.Print(m)
	m1 := map[int]string{111: "one", 222: "two", 333: "three"}
	fmt.Println(m1)
	for key, val := range m1 {
		fmt.Printf("%s => %d \n", key, val)
	}
	fmt.Println("\n------------map end------------")

}
func arrayDemo() {
	fmt.Println("\n------------array start------------")

	var array [5]int
	array[1] = 23
	array[3] = 78
	fmt.Println("定义时先不初使化", array)

	array1 := [5]int{1, 2, 3, 4, 5}
	fmt.Print("定义时初使化好", array1)
	fmt.Println("数组长度：", len(array1))
	b := array1[2:4] // a[2] 和 a[3]，但不包括a[4]
	fmt.Println("取数组[2,4)", b)

	b = array1[:4] // 从 a[0]到a[4]，但不包括a[4]
	fmt.Println("取数组[0,4)", b)

	b = array1[2:] // 从 a[2]到a[4]，且包括a[2]
	fmt.Println("取数组[2,4]", b)

	//多维数组
	var mutiArray [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			mutiArray[i][j] = i + j
		}
	}
	fmt.Println("二维数组: ", mutiArray)
	fmt.Println("\n------------array end------------")

}

func timeDemo() {
	fmt.Println("------------time start------------")
	var nowTime = time.Now()
	fmt.Println(nowTime)
	fmt.Println("\n------------time end------------")
}

func varDemo() {
	fmt.Println("------------var start------------")

	const hello1 string = "hello"
	const pi1 float32 = 3.1415926
	const (
		pi2    = 3.14
		hello2 = "world"
	)
	fmt.Println(hello1)
	fmt.Println(pi1)
	fmt.Println(hello2)
	fmt.Println(pi2)

	var vint1 = 12
	var world1 = "nihao"
	var (
		vint2  int    = 100
		world2 string = "字符串变量测试"
	)
	temp := "未先定义变量"
	fmt.Println(vint1)
	fmt.Println(world1)
	fmt.Println(vint2)
	fmt.Println(world2)
	fmt.Println(temp)
	fmt.Println("\n------------var end------------")

}

func forDemo() {
	fmt.Println("------------for start------------")
	//经典的for语句 init; condition; post
	for i := 0; i < 3; i++ {
		fmt.Print(i)
	}
	//精简的for语句 condition
	i := 1
	for i < 3 {
		fmt.Print(i)
		i++
	}
	//死循环的for语句 相当于for(;;)
	ii := 1
	for {
		if ii > 3 {
			break
		}
		fmt.Print(ii)
		ii++
	}
	fmt.Println("\n------------for end------------")
}

/**
*不定参数函数
 */
func sum(nums ...int) int {
	fmt.Print(nums)
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

/**
*返回两个结果的函数
 */
func multiReturn(key string) (int, bool) {
	m := map[string]int{"one": 1, "two": 2, "three": 3}

	var err bool
	var val int

	val, err = m[key]

	return val, err
}
