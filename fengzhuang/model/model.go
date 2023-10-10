package model
import "fmt"
type person struct{
	Name string
	age  int
    sal  float64
}
//写一个工厂模式的函数
func Newoerson (name string) *person {
	return &person{
		Name : name,
	}
}
//为了访问age和sal，编写一个get和set方法
//年龄的查看和设置
func ( p *person) SetAge(age int)  {
	if age  > 0 && age < 200 {
		p.age = age
	}else {
		fmt.Println("年龄不正确")
	}
}
	func (p *person) GetAge() int {
		return p.age
	}
//薪水的查看和设置
func ( p *person) SetSal(sal float64 ) {
	if sal > 0  {
		p.sal = sal
	}else {
		fmt.Println("薪水不对")
	}
}
func (p *person) GetSal() float64 {
	return p.sal
}