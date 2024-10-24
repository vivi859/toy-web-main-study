package main

import "fmt"

// 结构体测试,指针 slice map都是nil 没有分配空间，引用要make
type Cat struct {
	Name   string //=字段=属性=filed 可以是数组和引用类型，没有赋值就对应一个零值，print(*ptr.age)不能这样写.的优先级比＊高
	Age    int
	Colors string
	//test   [3]int            //数组
	ptr   *int              //指针 默认nil打印出就是nil
	slice []int             //切片 默认nil打印是空
	map1  map[string]string //切片 默认nil打印是空
}

func main() {
	//结构体测试
	var cat1 Cat
	cat1.Name = "小白"
	cat1.Age = 1
	cat1.Colors = "白色哦"
	cat1.slice = make([]int, 2)
	cat1.slice[0] = 22
	cat1.map1 = make(map[string]string) //这个大小不需要设置会自动增长
	cat1.map1["key1"] = "value12"

	cat2 := cat1 //值赋值　就是在内存中是另外一个内存空间，复制了一份cat内存和多一个变量cat2
	cat2.Colors = "黄色哦"

	cat3 := Cat{
		Name:   "小花花",
		Age:    1,
		Colors: "黄黄的",
	}

	var cat4 *Cat = new(Cat)
	var cat5 *Cat = &Cat{} //也可以这样定义

	cat4.Name = "指针猫咪" //*cat4.name=xx　标准应该这样写，go底层为了简化做了优化
	cat4.Age = 2
	cat4.Colors = "指针猫咪颜色花"

	fmt.Println(cat5)
	fmt.Println(cat4)
	fmt.Println(*cat4)
	fmt.Println(cat2)
	fmt.Println(cat3)
	fmt.Println(cat1)

}
