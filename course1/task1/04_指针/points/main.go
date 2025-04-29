package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int = 100
	// &a : 取变量 a 的内存地址
	var p *int = &a // p是指向a的指针

	fmt.Println(p)  // 打印地址
	fmt.Println(*p) // 取出p指向的值（就是a）

	*p = 200       // 修改p指向的值，相当于修改a
	fmt.Println(a) // 输出：200

	modifyPointer()

	unsafePointer()
}

func modifyPointer() {
	fmt.Println("start modify pointer")

	var p1 *int
	i := 1
	p1 = &i
	fmt.Println("*p1 == i ->", *p1 == i)
	*p1 = 2
	fmt.Println("i=", i)

	// 指针变量，同样和普通变量一样，会持有一个内存地址，内存地址对应的内存区域中保存的数据是另外一个变量的内存地址
	a := 2
	// 当使用*p 时，就是利用保存的这个内存地址访问到了实际的内存区域中的数据
	var p *int
	fmt.Println(&a)
	p = &a
	fmt.Println(p, &a)

	// **int 是指向 *int 的指针，指向指针的指针，叫做 二级指针（Pointer to Pointer）
	var pp **int = &p
	fmt.Println("pp, p:", pp, p)
	**pp = 3
	fmt.Println("pp, *pp, p :", pp, *pp, p)
	fmt.Println("**pp, *p:", **pp, *p)
	fmt.Println("a, &a:", a, &a)

}

func unsafePointer() {
	a := "Hello, world!"
	// 把a的指针转成uintptr
	upA := uintptr(unsafe.Pointer(&a))
	upA += 1

	// 重新强转回来，假如在下面这行代码之前发生了GC，*c可能会变成野指针，*c不能被打印
	c := (*uint8)(unsafe.Pointer(upA))
	fmt.Println(*c)
}
