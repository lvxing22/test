package main
// import "encoding/json"
import "fmt"
type student struct{
	Name  string
	Age   int
}

func (str *student) string()string{
	str.Name = "tot"
	lv := fmt.Sprintf("他的名字是=%v，他的年龄是=%v",str.Name,str.Age)
	return lv
}
func main (){
	//结构体之间可以强行转换，前提是字段得一样,个数，名称类型都得一样



	  str := student{
		Name :"tom",
		Age : 18,
	  }
	  fmt.Println(&str)
	 fmt.Println(str.string())
}