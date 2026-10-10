package simplebanksystem

type Bank struct {
	balance []int64
}

func Constructor(balance []int64) Bank {
	return Bank{
		balance: balance,
	}
}

func (this *Bank) isValidAccount(account int) bool {
	if account > len(this.balance) {
		return false
	}
	return true
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if !this.isValidAccount(account1) || !this.isValidAccount(account2) {
		return false
	}
	if !this.hasBalanceGreaterThan(account1, money) {
		return false
	}
	this.withdrawNoValidation(account1, money)
	this.depositNoValidation(account2, money)
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if !this.isValidAccount(account) {
		return false
	}
	this.depositNoValidation(account, money)
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if !this.isValidAccount(account) {
		return false
	}
	if !this.hasBalanceGreaterThan(account, money) {
		return false
	}
	this.withdrawNoValidation(account, money)
	return true
}

func (this *Bank) hasBalanceGreaterThan(account int, money int64) bool {
	return this.balance[account-1] >= money
}

func (this *Bank) withdrawNoValidation(account int, money int64) {
	this.balance[account-1] -= money
}

func (this *Bank) depositNoValidation(account int, money int64) {
	this.balance[account-1] += money
}

/**

 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
