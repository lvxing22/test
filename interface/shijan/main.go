package main

import (
	"fmt"
	"sort"
	"math/rand"
)
type Hos struct{
	Name string 
	Age  int
	Score int
}
type He []Hos
func (H He) Len() int{
    return len(H)
}
func (H He) Less(i,j int) bool{
	return H[i].Score < H[j].Score
}
func (H He) Swap(i,j int){
	H[i],H[j] = H[j],H[i]
}

  func main(){
	var her He
	for i :=0 ; i<10;i++{
		hos := Hos{
			Name :fmt.Sprintf("tom%d",rand.Intn(100)),
			Age : rand.Intn(100),
			Score :  rand.Intn(100),
		}
		her = append (her,hos)
	}
	//查看排序前的结果
	for _,v := range her{
		fmt.Println(v)
	}
	sort.Sort(her)
	fmt.Println("--------排序后的结果------")
	for _,v := range her{
		fmt.Println(v)
	}
	fmt.Println("",her)
}