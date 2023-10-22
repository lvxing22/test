package service
import (
	"example.com/TEST001/package01/custmermanage/model"
)
type Custmerservice struct{
	custmers []model.Custmer
	//表示当前有多少个客户，
	custmerNum int
}
func NewCustmerservice() *Custmerservice{
	custmerservice := &Custmerservice{}
	custmerservice.custmerNum = 1
	custmer := model.NewCustmer(1,"张山","男",22,"2646","zs@163.com")
	custmerservice.custmers = append(custmerservice.custmers,custmer)
	return custmerservice
}
//返回客户切片
func (this *Custmerservice) List() []model.Custmer{
	return this.custmers
}
func (this *Custmerservice) Add(custer model.Custmer) bool{
	this.custmerNum++
	custer.Id=this.custmerNum
	this.custmers = append(this.custmers,custer)
	return true
}
//根据id删除客户
func (this *Custmerservice) Delete(id int) bool{
	index := this.FindById(id)
	if index == -1{
		return false
	}
	//从切片中删除一个元素
	this.custmers = append(this.custmers[:index],this.custmers[index+1:]...)
	return true
}
// 根据id查找用户在切片中对应的下标,如果没找到该客户,返回-1
func (this *Custmerservice) FindById(id int) int {
index := -1
//遍历tish.custmers 切片
for i := 0; i < len(this.custmers); i++{
	if this.custmers[i].Id == id {
		index = i
	}
}
return index 
}