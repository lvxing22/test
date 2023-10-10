package main

import "fmt"
import "example.com/TEST001/fengzhuang/model"
//import "goproject/TEST001/fengzhuang/model"

//封装(encapsulation)就是把抽象出的字段和对字的操作封装在一起,数据被保护在内部,程序的其它包只有通过被授权的操作(方法),才能对字段进行操作。
//封装的理解和好处
// 隐藏实现细节
// 提可以对数据进行验证，保证安全合理
// 如何体现封装1) 对结构体中的属性进行封装2) 通过方法，包 实现封装
//案例：不能随便查看人的年龄,工资等隐私，并对输入的年龄进行合理的验证。
func main (){
	p := model.Newoerson("tom")
	fmt.Println(p)
	p.SetAge(13)
	p.SetSal(20000)
	fmt.Println(p)
}