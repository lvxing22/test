package main

import "fmt"
type Phone struct{

}
type Switch struct{

}
type Conpter struct{

}
type Usb interface {
	Start ()
	Stop ()
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
func (C Conpter) Working( usb Usb){
	usb.Start()
	usb.Stop()
}
func main (){
	phone := Phone{}
	swi   := Switch{}
	con   := Conpter{}
	con.Working(phone)
	con.Working(swi)
}