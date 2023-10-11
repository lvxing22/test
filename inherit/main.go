package main

import "fmt"
//解决代码冗余问题
// 继承可以解决代码复用,让我们的编程更加靠近人类思维
// 当多个结构体存在相同的属性(字段)和方法时,可以从这些结构体中抽象出结构体(比如刚才的student),在该结构体中定义这些相同的属性和方法
// 其它的结构体不需要重新定义这些属性和方法，只需嵌套一个Student匿名结构体即可。[画出示意图]
// 也就是说:在Golang中，如果一个struct嵌套了另一个匿名结构体，那么这个结构体可以直接访问匿名结构体的字段和方法，从而实现了继承特性。 、
/*
基本语法
type Goods struct {
	Name string
	Price float64
}
type Book struct {
	Goods
	Writer string
}
*/
//结构体和方法保存相同属性
type Student struct{
	Name string
	Age  int
	Score int
}
func (stu *Student) StuInto (){
	fmt.Printf("学生的名字=%v,学生的年纪=%v,学生的成绩=%v\n",stu.Name,stu.Age,stu.Score)
}
func (stu *Student) SetScore( score int){
	stu.Score = score
}
//不同的功能用不同的结构体和方法,并且镶嵌匿名结构体
type Grastudent struct{
	Student
}
func (p *Grastudent) testing (){
	fmt.Println("大学生在考试没成绩")
}
type Xxsstudent struct{
	Student
}
func (p *Xxsstudent) testing() {
	fmt.Println("小学生在考试没成绩")
}
//使用
func main(){
	xxs := &Xxsstudent{}
	xxs.Student.Name = "tom"
	xxs.Student.Age = 18
	xxs.testing()
	xxs.Student.Score = 99
	xxs.Student.StuInto()
//还可以这么使用
	Dxs := &Grastudent{}
	Dxs.Name = "roo"
	Dxs.Age = 30
	Dxs.testing()
	Dxs.Score = 70
	Dxs.StuInto()
}
