package main


import "fmt"
func main(){
	q1 := 3
	lvxing(q1)
	 res := lvxing(q1)
	 fmt.Println(res)
}
func  lvxing(q1 int) string{
	avc := [5]int{2,3,4,5,7}
	 for j := 0 ;j <len(avc);j++{
		 if q1 == avc[j]{
			return "在中间"
		 }
}
return"不在中间"
}