package main

import "fmt"

func test (){
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println("err=",err)
		}
	}()
	num1 := 10
	num2 := 0
	num3 := num1 / num2
	fmt.Println("num3=",num3)
	
}
func main (){
	test()
	fmt.Println("main")
}