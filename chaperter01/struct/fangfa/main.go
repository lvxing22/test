package main
import "fmt"
type Ione struct{

}
func (c Ione) Met(len int,wit int)int{
	return len * wit 
}
func main(){
	var c Ione

	fmt.Println(c.Met(23,32))
}