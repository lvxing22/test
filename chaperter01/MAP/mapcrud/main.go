package main


import "fmt"
func main (){
    //增加修改，key存在就修改，不存在就增加
	var a map [int] int
	a = make(  map [int] int,10)
	a[3]=3
	a[2]=888
	a[4]=999
	a[10]=88
	fmt.Println(a)
	a[2]=4
	a[11]=100
	fmt.Println(a)
}