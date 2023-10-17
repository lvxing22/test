package main
import (
	"fmt"
)
//显示主菜单的、
func main(){
	//声明一个变量接受key，保存用户输入的选项
	key := ""
	//声明一个变量，接收推出软件
	Loop := true
	//声明账户余额
	balance := 10000.00
	//声明收支
	money := 0.00
	//声明收支
	//收支说明
	 note := ""
	//声明说明
	//收支的详细用字符串表示
	ditails :=("收支\t账户金额\t收支金额\t说明")
	//声明是否有收支情况
	flag := false
	
	//显示主菜单
   for{

	fmt.Println("\n------------------家庭收支记账软件-------------------")
	fmt.Println("                   1 收支明细")
	fmt.Println("                   2 登记收入")
	fmt.Println("                   3 登记支出")
	fmt.Println("                   4 推出登录")
	fmt.Println("请选择1-4: ")
	fmt.Scanln(&key)
	switch key {
	case "1":
		if flag{
			fmt.Println("---------------当前收支明细记录---------------")
			fmt.Println(ditails)
		
		}else{
			fmt.Printf("目前总金额为10000元但是没有收支情况。。。")
		}
	case "2":
		flag = true
		fmt.Printf("输入本次收入的金额：\n")
		fmt.Scanln(&money)
		balance += money
		fmt.Printf("收入说明：\n")
		fmt.Scanln(&note)	
		fmt.Printf("本次收入金额：%v\n",money)
		fmt.Printf("本次收入说明：%v\n",note)
		ditails += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v",balance,money,note)
	case "3":	
	flag = true
	fmt.Printf("输入本次支出的金额：\n")
	fmt.Scanln(&money)
	if money > balance {
		fmt.Println("钱不够了")
		break
	}
	balance -= money
	fmt.Printf("支出说明：\n")
	fmt.Scanln(&note)	
	fmt.Printf("本次支出金额：%v\n",money)
	fmt.Printf("本次支出说明：%v\n",note)
	ditails += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v",balance,money,note)
	case "4":	
		yes :=""
	    fmt.Println("确定要推出吗？[Y/n]")
		for {
		fmt.Scanln(&yes)
			if yes == "Y"{
			 Loop = false
			return
			}else if yes == "n"{
				break
			}else {
				fmt.Println("请输入真确的值[Y/n]")
			}
		}
	default :
	    fmt.Println("请输入正确的选项")
	if !Loop {
		break
	}

	}
   }
}