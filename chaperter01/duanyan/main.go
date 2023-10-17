package main

import "fmt"
type Point struct{
	Name int
}
func main(){
	var a interface{}
	var  point Point = Point{ 12}
	a = point
    var b Point
	 b = a.(Point)
	 fmt.Println(b) 
//断言带检测
var  a1 interface{}
   var po  float64 = 33
   a1 = po
     b1,ok  := a1.(float64)
	 if ok == true {
		fmt.Printf("转换成功最终的值为= %v 类型为%T",b1,b1)
	}else {
		fmt.Println("转换失败")
	}
}