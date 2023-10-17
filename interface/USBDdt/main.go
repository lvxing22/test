package main

import "fmt"
type Phone struct{
    Name string
}
type Switch struct{
     Name  int
}
type Conpter struct{

}
type Usb interface {
	
}
func (p Phone) Start (){
	fmt.Println("手机开始工作了。。。")
}
func (p Phone) Stop (){
	fmt.Println("手机停止工作了。。。")
}
func (S Switch) Start (){
	fmt.Println("游戏机开始工作了。。。")
}
func (S Switch) Stop (){
	fmt.Println("游戏机停止工作了。。。")
}
// func (C Conpter) Working( usb Usb){
// 	usb.Start()
// 	usb.Stop()
// }
func main (){
	var A [2]Usb
	A[0] = Phone{Name:"tom"}
	A[1] = Switch{Name: 12}
	fmt.Println(A)
}