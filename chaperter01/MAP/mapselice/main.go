package main


import "fmt"
func main(){
	//selice map 可以使map动态增长
	 var mon  []map[string]string
	mon = make([]map[string]string,2)
	if mon[0] ==nil{
		mon[0] = make(map[string]string,2)
        mon[0] ["name"]="孙悟空"
		mon[0] ["age"]="100"

	fmt.Println(mon)
}
if mon[1] ==nil{
	mon[1] = make(map[string]string,2)
	mon[1] ["name"]="猪八戒"
	mon[1] ["age"]="10"

fmt.Println(mon)
   //动态增加
   newmon := map[string]string{
	  "name":"狮子精","age":"1000",
   }
   mon = append(mon,newmon)
   fmt.Println(mon)
}
}