package main

import "fmt"

func test (){
	defer func(){
		err := recover()
		if err != nil{
			fmt.Println("err=",err)
		}
	}()
	a := 10
	b := 0
	c := a / b
	fmt.Println("c=",c)
}
func main (){
	test()
	fmt.Println("testetesat")
}