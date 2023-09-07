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
	//删,如果k不存在也不会报错
	delete(a,3)
	fmt.Println(a)
	delete(a,5)
	fmt.Println(a)
	val,ok :=a[3]
	if ok {
		fmt.Println("存在",val)
	}else{
		fmt.Println("不存在",val)
	}
	//删除所有的key第一种办法重新make一次
	a = make( map [int] int,10)
	fmt.Println(a)
	//查询
	a[0]= 4
    a[1]= 3
	a[3]=99
	_,ok2 :=a[3]
	if ok2 {
		fmt.Println("存在")
	}else{
		fmt.Println("不存在")
	}
	//map的便利
	for k,v := range a{
		fmt.Printf("k=%v,v=%v\n",k,v)
	}
	//多层遍历
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
	for k1,k2 := range exec{
		for k2,v1 :=range k2{
			fmt.Printf("k1=%v k2=%v v1=%v\n",k1,k2,v1)
		}
	}
//统计map有多少对kv
 fmt.Println(len(a))
}