package model
import(
	"fmt"
)
type Custmer struct{
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}
func NewCustmer(id int, name string, gender string, age int, phone string, email string) Custmer{
	return Custmer{
		Id : id,
		Name : name,
		Gender : gender,
		Age : age,
		Phone : phone,
		Email : email,
	}
}
func NewCustmer2(name string, gender string, age int, phone string, email string) Custmer{
	return Custmer{
		Name : name,
		Gender : gender,
		Age : age,
		Phone : phone,
		Email : email,
	}
}
//创建第二种Custer实例的方法,这个方法不带id

func (this Custmer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v\t",this.Id,this.Name,this.Gender,this.Age,this.Phone,this.Email,)
	return info
}

