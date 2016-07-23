package main

import (
	"fmt"
	"time"
)

func main() {

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

	//时间相关
	var nowTime = time.Now()
	fmt.Println(nowTime)
	//数组
	var array [5]int
	array[1] = 23
	array[3] = 78
	fmt.Println("定义时先不初使化", array)

	array1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("定义时初使化好", array1)
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

	//函数测试，有一个返回值
	fmt.Println("求和函数", sum(12, 11))

	forDemo()

}

func forDemo() {
	//经典的for语句 init; condition; post
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	//精简的for语句 condition
	i := 1
	for i < 3 {
		fmt.Println(i)
		i++
	}
	//死循环的for语句 相当于for(;;)
	ii := 1
	for {
		if ii > 3 {
			break
		}
		fmt.Println(ii)
		ii++
	}
}

func sum(a int, b int) int {
	return a + b
}
