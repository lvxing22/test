package main

import "fmt"
func main(){
	//切片的三种声明方式
	//第一种
	var a1 =[3] int { 1,2,3}
	a2 := a1[0:1]
	fmt.Println(a2)
	//第二种
	 var b1 []int = make ( []int, 5, 10)
	 fmt.Println(b1)
	// b1 [1] =33
	// b1 [2] =44 
	var b2 [] string =make([]string,4,6)
	b2[1]="tom"
	fmt.Println(b2)
	fmt.Println("b2的容量=",cap(b2))
    fmt.Println("b2的长度=",len(b2))
	 //第三种
    var c1 [] int = []int {1,2,3}
	fmt.Println(c1)
	fmt.Println("c1的容量=",cap(c1))
    fmt.Println("c1的长度=",len(c1))
	//切片的两种遍历方式
	//第一种
	for i := 0 ;i < len(c1);i++{
		fmt.Printf("c1[%v]=%v\n",i,c1[i])
	}
	//第二种
	for z,x :=  range b2 {
		fmt.Printf("b2[%v]=%v\n",z,x)
	//append 动态增长切片 
	var d1 [] int = []int {1,2,3}
	fmt.Println(d1)
	fmt.Println("d1的容量=",cap(d1))
    fmt.Println("d1的长度=",len(d1))
	 d1 =append (d1 ,333,444,555)
	fmt.Println(d1)
	fmt.Println("d1的容量=",cap(d1))
    fmt.Println("d1的长度=",len(d1))
    //append 切片增长切片
	d1 =append (d1 ,c1...)
	fmt.Println(d1)
	fmt.Println("d1的容量=",cap(d1))
    fmt.Println("d1的长度=",len(d1))
	//切片的拷贝 *如果容量不够也不会报错，只是不会拷贝全

	e1 := [] int {1,2,3}
	var e2 [] int = make ([] int ,6,7)
	 copy (e2,e1 )
	fmt.Println(e1)
	fmt.Println(e2)
	}
     
	lef := [] int {1,2,3}
	zongshu := 0
	for i :=0 ;  i < len(lef);i ++{
	 zongshu += lef[i]
}
	 fmt.Println(zongshu)
}