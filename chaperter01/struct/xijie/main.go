package main

import "fmt"
func main (){
	//结构体之间可以强行转换，前提是字段得一样,个数，名称类型都得一样
	type a struct{
		num int
	}
	type b struct{
		num int
	}
	var  A a
	var B  b
	A = a(B)
	fmt.Println(A,B)
}