package main

import "fmt"
import "example.com/TEST001/chaperter01/struct/gcms/model"
func main()	{
var name = model.Student{
		Name : "tom",
	}
	fmt.Printf(name.Name)
	fmt.Println()
	n := "tooom~"
	model.Newstudent(n)
	fmt.Println(*model.Newstudent(n))
}