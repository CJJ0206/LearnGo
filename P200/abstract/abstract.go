package abstract

import "fmt"

/*
8.13
抽象
	我们在前面定义一个结构体的时候，实际上九十八一类事物的共有属性（字段）和行为（方法）提取出来，幸好曾一个物理模板。
	这种研究问题的方法我们称为抽象
*/

type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

// Deposit 存款
func (account *Account) Deposit(money float64, Pwd string) {
	if Pwd != account.Pwd {
		fmt.Println("你输入的密码不正确！")
		return
	}
	if money < 0 {
		fmt.Println("金额不正确！")
		return
	}
	account.Balance += money
	fmt.Println("存款成功。")
}

// WithDraw 取款
func (account *Account) WithDraw(money float64, Pwd string) {
	if Pwd != account.Pwd {
		fmt.Println("你输入的密码不正确！")
		return
	}
	if money <= 0 {
		fmt.Println("账户金额不够！")
		return
	}
	account.Balance -= money
	fmt.Println("取款成功。")
}

// GetBalance 查询余额
func (account *Account) GetBalance(pwd string) {
	if account.Pwd != pwd {
		fmt.Println("密码不正确")
		return
	}
	fmt.Println(account.Balance)
}

// 话可以做的设计就是让用户在终端自己出入想要做的事情然后执行（菜单）
