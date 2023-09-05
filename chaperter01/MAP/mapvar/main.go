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
	//第二种声明方式
	b := make(map [int] int,10)
	b[2]=4
	fmt.Println(b)
	////第三种声明方式
	c := map [int] int{
		1:3,
		2:4,
	}
	fmt.Println(c)
	exec := make(map[int]map[string]string)
	exec[1]=make(map[string]string,3)
	exec[1]["性别"]="男"
	exec[1]["年龄"]="18"
	exec[1]["家庭地址"]="北京"
	exec[2]=make(map[string]string,3)
	exec[2]["性别"]="女"
	exec[2]["年龄"]="10"
	exec[2]["家庭地址"]="上海"
	fmt.Println(exec)
	fmt.Println(exec[2]["性别"])
	
}