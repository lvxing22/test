package main


import "fmt"
func main (){
	//第一种声明方式
	var a map [int] int
	a = make(  map [int] int,10)
	a[3]=3
	a[2]=888
	a[3]=999
	a[10]=88
	fmt.Println(a)

}