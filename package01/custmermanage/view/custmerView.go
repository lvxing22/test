package main
import(
	"fmt"
	"example.com/TEST001/package01/custmermanage/service"
	"example.com/TEST001/package01/custmermanage/model"
)
type custmerView struct{
	key string// 接收用户输入
	loop bool// 判断是否退出程序
	// 增加一个custmerService这个字段
	custmerService *service.Custmerservice
}
// 显示所有客户信息
func (this *custmerView) List(){
	custers := this.custmerService.List()
	fmt.Println("--------------客户列表--------------")
	fmt.Printf("编号\t姓名\t性别\t年龄\t电话\t邮箱\n")
	for i := 0 ; i < len(custers); i++{
		fmt.Println(custers[i].GetInfo())
	}
	fmt.Printf("--------------客户列表完成--------------\n")
}
// 得到用户的输入，信息构建新的客户，并完成添加
func (this *custmerView) add(){
	fmt.Println("--------------添加客户--------------")
	fmt.Println("姓名:")
	name :=""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone :=""
	fmt.Scanln(&phone)
	fmt.Println("邮箱:")
	email :=""
	fmt.Scanln(&email)

	// 构建一个新的custmer
	custmer :=model.NewCustmer2(name,gender,age,phone,email)
	// 调用
	if this.custmerService.Add(custmer) {
		fmt.Println("--------------添加完成--------------")
	} else{
		fmt.Println("--------------添加失败--------------")
	}
}
	//输入用户id,删除该id对应的用户
	func (this *custmerView) delete(){
		fmt.Println("--------------删除用户--------------")
		fmt.Println("请输入待删除用户的编号(-1退出): ")
		id := -1
		fmt.Scanln(&id)
		if id == -1 {
			return
		}
		fmt.Println("确认是否删除(Y/N):")
		choice :=""
		fmt.Scanln(&choice)
		if choice == "Y" || choice =="y"{
			if this.custmerService.Delete(id) {
				fmt.Println("--------------删除完成 --------------")
			}else{
				fmt.Println("--------------输入的id号不存在 --------------")
			}
		}

	}
func (this *custmerView) mainView(){
	for{

		fmt.Println("--------------客户信息管理软件--------------")
		fmt.Println("               1 添加客户                  ")
		fmt.Println("               2 修改客户                  ")
		fmt.Println("               3 删除用户                  ")
		fmt.Println("               4 客户列表                  ")
		fmt.Println("               5 退出                      ")
		fmt.Println("               请选择 [1-5]                ")

		fmt.Scanln(&this.key)
		switch this.key {
			case "1" :
				this.add()
			case "2" :
				fmt.Println("               2 修改客户                  ")
			case "3" :
				this.delete()
			case "4" :
				fmt.Println("               4 客户列表                  ")
				this.List()
			case "5" :
				this.loop = false
			default :
				fmt.Println("               输入有误重新输入                 ")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("退出成功~")
}
func main() {
	custmerView :=custmerView{
		key : "",
		loop : true,
	}
	// 对custmerView中的custmerService初始化
	custmerView.custmerService = service.NewCustmerservice()
	// 显示菜单
	custmerView.mainView()
}

