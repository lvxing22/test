package main

import "fmt"
var (
	n1 = func ( n1,n2 int)int{
		return n1 * n2 
	}
)
func  test() func (int) int{
	 var a  = 100
	return func (x int) int {
		a = a +x 
		return a
	} 
	 
}
func main (){
	f := test()
	fmt.Println(f(1))
	 res:= func (n1 int,n2 int) int{
		return n1 + n2 
	}(3333,3333)
	fmt.Println("res=",res) 
   a := func ( n2 int, n1 int)int{
          return n2 + n1
   } 
   a1 := a(2222,3333)
   a2 := a(22222,9999)
   fmt.Printf("a1= %v,a2= %v\n",a1,a2 )
   fmt.Println("a2= ",a2 )
   a3 := n1(222,333)
   fmt.Println("a3=",a3)
   var s1 [3] string
   s1 [0] = "2"
   s1 [1] = "1"
   s1 [2] = "3"
   fmt.Println("s1",s1)
   q1 := "12332313"
   fmt.Println("q1=",len(q1))
  }

