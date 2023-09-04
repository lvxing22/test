package main

import "fmt"

func main(){
	 //交换排序 1、冒泡排序，快速排序
	 arr := [5] int{4,6,11,7,5}
	 fmt.Println(arr)
	 test(&arr)
	fmt.Println(arr)
	q1 := 11
   fmt.Println(lvxing(q1))
}
 func test( arr *[5]int) {
	for i :=0; i <4; i++{
	 for j :=0 ;j < len(*arr)-1-i; j++{
		if (*arr)[j] > (*arr)[j+1]{
			temm := 0
			temm  =(*arr)[j]
			(*arr)[j]=(*arr)[j+1]
			(*arr)[j+1] = temm
		}
		fmt.Println("第%v轮地%v次排序的结果=%v",i+1,j+1,*arr)
	}
}
 }
    //查找1、顺序查找 2、二分查找
	func  lvxing(q1 int) string{
		avc := [5]int{2,3,4,5,7}
         for j := 0 ;j <len(avc);j++{
			 if q1 == avc[j]{
				return "在中间"
			 }
		 }
		 return "不在中间"
		}
     //二分查找
	