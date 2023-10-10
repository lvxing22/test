package main

import "fmt"
//写一个银行的小功能，银行有 账户，密码，名称 可以查询，存钱，取钱
//定义一张银行卡的结构体
type Account struct{
	Account1 string
	Pwd      string
	Balance  float64
}
//定义一个可以存款的方法
func (A *Account)  Deposite(money float64 ,pwd string){
	//输入的密码是否正确
	if pwd != A.Pwd {
		fmt.Println("密码不正确")
		return
	} 
	//输入的存款金额是否正确
	if  money  <= 0.00 {
		fmt.Println("金额不正确")
		return
	}
	A.Balance += money
	fmt.Println("存款成功")
	fmt.Printf("一共还有%v元",A.Balance)
}
//定义一个可以取款的方法
func (A *Account)  Withdraw(money float64 ,pwd string){
	//输入的密码是否正确
	if pwd != A.Pwd {
		fmt.Println("密码不正确")
		return
	} 
	//输入取款金额是否正确
	if  money > A.Balance || money <= 0 {
		fmt.Println("余额不足")
		return
	}
	A.Balance -= money
	fmt.Println("取款成功")
	fmt.Printf("一共还有%v元",A.Balance)
}
//定义一个可以查询余额的方法
func (A *Account)  Query(pwd string){
	//输入的密码是否正确
	if pwd != A.Pwd {
		fmt.Println("密码不正确")
		return
	} 
	fmt.Printf("一共还有%v元\n",A.Balance)
}
func main (){
	account := Account{
			Account1 : "gs",
			Pwd      :  "666666",
			Balance  : 100.00,
	}
	account.Query("666666")
	account.Deposite(3222,"666666")
	account.Withdraw(3222,"666666")
}