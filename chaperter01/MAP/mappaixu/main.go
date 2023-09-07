package main


import "fmt"
//import "sort"
func main (){
	//将map放入selice中对celice进行排序再遍历切片，然后按照key输出map
    var mon  map[string]string
	mon =make (map[string]string)
	mon["wqe"] ="123"
	mon["321"]="321"
	mon["qwee"]="1233221"
	mon["ewqewqe"]="asdwa"
	//mon[3]="wqe2"
	//mon[4]="dsf2"
	//mon[43]="sdwawr" 
	fmt.Println(mon)
	
}
