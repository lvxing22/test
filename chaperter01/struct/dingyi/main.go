package main

import "fmt"
func  main (){
	//定义一个猫的结构体
    /*
	  声明方式：
	  tpye  名称 struct{
		a int
		b  string
	  }
	  1、如果名称的第一个首字母大写代表可以在其他包使用
	  2、如果字段的第一个首字母大写代表可以在其他包引用
	  3、字段的类型可以是基本数据类型也可以是引用类型
	*/
	type cat struct {
		name string
		age  int
		color string
	}
	var cat1 cat
	fmt.Println(cat1)
	cat1.name = "小白"
	cat1.age = 2
	cat1.color = "黑色"
	fmt.Println(cat1)
	fmt.Println("输出猫的信息：")
	fmt.Printf("猫的名字=%v\n",cat1.name)
	fmt.Printf("猫的年龄=%v\n",cat1.age)
	fmt.Printf("猫的颜色=%v\n",cat1.color)
	/*type person struct{
		p1 []int
		//p2 *int
		p3 map[string]int
	}
/*	var pp person
	fmt.Println(pp)
	pp.p1 = make ( [] int ,2)
	pp.p1[0]=1
	pp.p3 =make(map[string]int)
	pp.p3["小明"]=18
	
	fmt.Println(pp)*/
	//第二种声明方法
	type lv struct {
		lq int 
		lw string
	}
	qq := lv{13,"23"}
    fmt.Println(qq)
	type person struct {
		p1 []int
		//p2 *int
		p3 map[string]int
	}
	ww := person{p1: make([]int, 2), p3: map[string]int{"int": 12}}
	fmt.Println(ww)
	//第三种声明方式
	var q1 *lv = new(lv)
	(*q1).lq = 123
	(*q1).lw = "2321"
	fmt.Println(*q1)
	fmt.Println(q1)
	//还可以不需要*
	q1.lq =321
	fmt.Println(*q1)
	fmt.Println(q1)
	//第四种声明方式
	var lr *lv = &lv{998,"nishizhu"} 
	fmt.Println(*lr)
}