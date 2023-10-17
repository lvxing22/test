package main

import (
	"fmt"
	_ "sort"
	_ "math/rand"
)
type mon struct{
	Name string
}
func (m *mon) pashu(){
	fmt.Println(m.Name,"猴子会爬树")
}
type mons struct{
	mon
}
// type fash interface{
// 	fashing()
// }
func (m *mons) fashing(){
	fmt.Println(m.Name,"猴子还会飞")
}
func main(){
	tom := mons{
		mon{
		Name : "TES",
		},
	}
	tom.pashu()
	tom.fashing()
}