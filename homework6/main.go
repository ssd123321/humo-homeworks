package main

import "fmt"

func main() {
	// 1
	balance0 := 100000
	monthly_addition := 1000

	for i := 0; i < 12; i++ {
		add_monthly_balance(&balance0, monthly_addition)
	}
	// 2
	credit_sum := 3000000
	interest_rate0 := 0.1

	for i := 0; i < 12; i++ {
		calculate_payment(&credit_sum, interest_rate0)
	}
	// 3
	balance1 := 500000
	transfer_limit := 100000

	for i := 0; i < 10; i++ {
		fmt.Printf("Введите сумму %d-го перевода: ", i+1)
		var amount int
		fmt.Scan(&amount)
		make_transfer(&balance1, transfer_limit, amount)
	}
	// 4
	deposit_sum0 := 100000
	interest_rate1 := 0.05

	for i := 0; i < 12; i++ {
		calculate_deposit(&deposit_sum0, interest_rate1)
	}
	// 5
	usd_rate := 74.5

	for i := 0; i < 12; i++ {
		update_rate(&usd_rate)
		fmt.Print("Введите сумму для конвертации: ")
		var amount float64
		fmt.Scan(&amount)
		convert_currency(usd_rate, amount)
	}
	// 6
	deposit_sum1 := 200000
	interest_rate := 0.05

	for i := 0; i < 12; i++ {
		calculate_deposit_compound(&deposit_sum1, interest_rate)
	}
	// 7
	daily_limit := 5000

	for i := 0; i < 10; i++ {
		fmt.Printf("Введите сумму %d-го расхода: ", i+1)
		var expense int
		fmt.Scan(&expense)
		add_expense(&daily_limit, &expense)
	}
	// 8
	deposit_sum2 := 500000
	interest_rate2 := 0.06

	for i := 0; i < 5; i++ {
		calculate_deposit_yearly(&deposit_sum2, interest_rate2)
	}
	// 9
	balance2 := 100000
	commission_rate := 0.01

	for i := 0; i < 10; i++ {
		fmt.Printf("Введите сумму %d-й транзакции: ", i+1)
		var amount int
		fmt.Scan(&amount)
		add_transaction(&balance2, commission_rate, &amount)
	}
	// 10
	investment_sum := 300000
	interest_rate3 := 0.07

	for i := 0; i < 5; i++ {
		calculate_investment_yearly(&investment_sum, interest_rate3)
	}
	// 11
	balance4 := 10000

	for i := 0; i < 5; i++ {
		fmt.Printf("Введите сумму %d-го пополнения: ", i+1)
		var amount int
		fmt.Scan(&amount)
		add_balance(&balance4, &amount)
	}
	// 12
	balance5 := 2000000

	for i := 0; i < 5; i++ {
		fmt.Printf("Введите сумму %d-го снятия: ", i+1)
		var amount int
		fmt.Scan(&amount)
		withdraw_balance(&balance5, &amount)
	}
	// 13
	balance6 := 150000
	check_balance(&balance6)
	// 14
	balance := 5000000
	interest_rate4 := 0.02
	calculate_interest(&balance, interest_rate4)
	// 15
	account1_balance := 3000000
	account2_balance := 1500000
	transfer_amount := 1000000
	transfer_funds(&account1_balance, &account2_balance, &transfer_amount)
	// 16
	balance7 := 4000000
	withdrawal_limit := 100000

	for i := 0; i < 5; i++ {
		fmt.Printf("Введите сумму %d-го снятия: ", i+1)
		var amount int
		fmt.Scan(&amount)
		withdraw_money(&balance7, withdrawal_limit, &amount)
	}
}
