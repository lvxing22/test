package utils
import "fmt"
type Ziduan struct{
	//声明一个变量接受key，保存用户输入的选项
	key  string
	//声明一个变量，接收推出软件
	Loop bool
	//声明账户余额
	balance float64
	//声明收支
	money float64
	//声明收支
	//收支说明
	 note string
	//声明说明
	//收支的详细用字符串表示
	ditails string
	//声明是否有收支情况
	flag bool
}
func Csh () *Ziduan {
    return &Ziduan{
		key : "",
		Loop : true,
		balance : 10000.00,
		money : 0.00,
		note : "",
		ditails : "收支\t账户金额\t收支金额\t说明",
		flag : false,
	}
}
func (this *Ziduan) Case1 (){
	if this.flag{
		fmt.Println("---------------当前收支明细记录---------------")
		fmt.Println(this.ditails)
	
	}else{
		fmt.Printf("目前总金额为10000元但是没有收支情况。。。")
	}
}
func (this *Ziduan)Case2(){
			this.flag = true
			fmt.Printf("输入本次收入的金额：\n")
			fmt.Scanln(&this.money)
			this.balance += this.money
			fmt.Printf("收入说明：\n")
			fmt.Scanln(&this.note)	
			fmt.Printf("本次收入金额：%v\n",this.money)
			fmt.Printf("本次收入说明：%v\n",this.note)
			this.ditails += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v",this.balance,this.money,this.note)
}
func (this *Ziduan)Case3(){
		this.flag = true
		fmt.Printf("输入本次支出的金额：\n")
		fmt.Scanln(&this.money)
		if this.money > this.balance {
			fmt.Println("钱不够了")
		}
		this.balance -= this.money
		fmt.Printf("支出说明：\n")
		fmt.Scanln(&this.note)	
		fmt.Printf("本次支出金额：%v\n",this.money)
		fmt.Printf("本次支出说明：%v\n",this.note)
		this.ditails += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v",this.balance,this.money,this.note)
}
func(this *Ziduan)Case4(){
			yes :=""
			fmt.Println("确定要推出吗？[Y/n]")
			for {
			fmt.Scanln(&yes)
				if yes == "Y"{
				 	this.Loop = false
					return
				}else if yes == "n"{
					break
				}else {
					fmt.Println("请输入真确的值[Y/n]")
				}
			}
}	
func (this *Ziduan) Zcd(){
	for{
		fmt.Println("\n------------------家庭收支记账软件-------------------")
		fmt.Println("                   1 收支明细")
		fmt.Println("                   2 登记收入")
		fmt.Println("                   3 登记支出")
		fmt.Println("                   4 推出登录")
		fmt.Println("请选择1-4: ")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.Case1()
		case "2":
			this.Case2()
		case "3":	
		this.Case3()
		case "4":	
			this.Case4()
			fmt.Println(this.Loop)
		default :
			fmt.Println("请输入正确的选项")
		// if !this.Loop {
		// 	fmt.Println("testtest")
		// 	break
		// }
		}
		if !this.Loop {
			fmt.Println("testtest")
			break
		}
	   
	}
}