package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

var pl = fmt.Println

type Account struct {
	username      string
	password      string
	balance       float64
	accountType   string
	accountNumber int
}

// var accounts []Account // Slice of accounts from previous implementation
var accounts map[int]Account
var accountsByUsername map[string]Account

func main() {
	accounts = make(map[int]Account)
	accountsByUsername = make(map[string]Account)
	reader := bufio.NewReader(os.Stdin)
	pl("==Welcome to Golangking System!==")
	loginScreen(reader)
}

func mainMenu(reader *bufio.Reader, account Account) {
	for {
		pl("1. Check Balance  2. Deposit  3. Withdraw  4. Fund Transfer  5. Logout")
		pl("Enter your choice : ")
		choice, err := reader.ReadByte()
		if err != nil {
			pl("Invalid input!\nPlease try again!")
			continue
		}

		reader.ReadString('\n')

		switch choice {
		case '1':
			checkBalance(account)
		case '2':
			deposit(reader, account)
		case '3':
			withdraw(reader, account)
		case '4':
			fundTransfer(reader, account)
		case '5':
			pl("Logging out...")
			loginScreen(reader)
			return
		default:
			pl("Invalid input!\nPlease try again!")
			continue
		}
	}
}

func deposit(reader *bufio.Reader, account Account) {
	for {
		pl("Enter the amount to deposit or -1 to cancel : ")
		input, err := reader.ReadString('\n')

		input = input[:len(input)-1]

		if err != nil {
			pl("Invalid input!\nPlease try again!")
			continue
		}

		if input == "-1" {
			pl("Deposit cancelled. Returning to main menu...")
			reader.ReadString('\n')
			return
		}

		amount, err := strconv.ParseFloat(input, 64)

		if err != nil {
			pl("Invalid input!\nPlease try again!")
			continue
		} else if amount <= 0 {
			pl("You cannot deposit a negative or zero amount!\nPlease try again!")
			continue
		} else if amount > 100000 {
			pl("You cannot deposit more than 100000 at a time!\nPlease try again!")
			continue
		}

		mappedAccount := accounts[account.accountNumber]
		mappedAccount.balance += amount

		accounts[account.accountNumber] = mappedAccount
		accountsByUsername[account.username] = mappedAccount

		pl("Deposit successful!")
		pl("Your new balance is : ", mappedAccount.balance)
		return
	}
}

func fundTransfer(reader *bufio.Reader, account Account) {
	userAccount := accounts[account.accountNumber]

	for {
		pl("Enter the account number of the recipient or -1 to cancel: ")
		input, err := reader.ReadString('\n')

		input = input[:len(input)-1]

		if err != nil {
			pl("Invalid input!\nPlease try again!")
			continue
		}

		if input == "-1" {
			pl("Fund transfer cancelled. Returning to main menu...")
			return
		}

		recipientAccountNumber, err := strconv.Atoi(input)

		if err != nil {
			pl("Invalid input!\nPlease try again!")
			continue
		}

		pl("Enter the amount to transfer : ")
		input, err = reader.ReadString('\n')

		input = input[:len(input)-1]

		if err != nil {
			pl("Invalid input")
		}

		amount, err := strconv.ParseFloat(input, 64)

		if err != nil {
			pl("Invalid input!\nPlease try again!")
			continue
		}

		if amount <= 0 {
			pl("You cannot transfer a negative or zero amount!\nPlease try again!")
			continue
		} else if amount > userAccount.balance {
			pl("Insufficient funds!\nPlease try again!")
			continue
		}

		recipientAccount, ok := accounts[recipientAccountNumber]

		if !ok {
			pl("Recipient account not found!\nPlease try again!")
			continue
		}

		recipientAccount.balance += amount
		userAccount.balance -= amount

		accounts[recipientAccountNumber] = recipientAccount
		accounts[userAccount.accountNumber] = userAccount
		accountsByUsername[userAccount.username] = userAccount
		pl("Account not found!\nPlease try again!")
	}
}

func withdraw(reader *bufio.Reader, account Account) {
	userAccount := accounts[account.accountNumber]

	for {
		pl("Enter the amount to withdraw or -1 to cancel : ")
		input, err := reader.ReadString('\n')

		input = input[:len(input)-1]

		if err != nil {
			pl("Invalid input!\nPlease try again!")
			continue
		}

		if input == "-1" {
			pl("Withdrawal cancelled!\nReturning to main menu...")
			reader.ReadString('\n')
			return
		}

		amount, err := strconv.ParseFloat(input, 64)

		if err != nil {
			pl("Invalid input!\nPlease try again!")
			continue
		}

		if amount <= 0 {
			pl("You cannot withdraw a negative or zero amount!\nPlease try again!")
			continue
		} else if amount > userAccount.balance {
			pl("Insufficient funds!\nPlease try again!")
			continue
		}

		userAccount.balance -= amount
		accounts[account.accountNumber] = userAccount
		pl("Withdrawal successful!")
		pl("Your new balance is : ", userAccount.balance)
	}
}

func checkBalance(account Account) {
	pl("Your balance is : ", account.balance)
}

func loginScreen(reader *bufio.Reader) {
	for {
		pl("\n1. Login  2. Register  3. Exit")
		pl("Enter you choice: ")
		choice, err := reader.ReadString('\n')

		choice = choice[:len(choice)-1]

		if err != nil {
			pl("Invalid input")
			continue
		}

		switch choice {
		case "1":
			login(reader)
		case "2":
			register(reader)
		case "3":
			pl("Exiting...")
			return
		default:
			pl("Invalid input")
		}
	}
}

func login(reader *bufio.Reader) {
	for {
		pl("Login")
		pl("Enter your username or -1 to cancel : ")
		username, err := reader.ReadString('\n')

		username = username[:len(username)-1]

		if err != nil {
			pl("Invalid input", err)
			continue
		}

		if username == "-1" {
			pl("Login cancelled!\nReturning to login screen...")
			return
		}

		pl("Enter your password: ")
		password, err := reader.ReadString('\n')

		password = password[:len(password)-1]

		if err != nil {
			pl("Invalid input", err)
			continue
		}

		value, ok := accountsByUsername[username]

		if !ok {
			pl("Account not found!\nPlease try again!")
			continue
		}

		if value.password != password {
			pl("Incorrect password!\nPlease try again!")
			continue
		}

		pl("Login successful!")
		pl("Welcome ", username)
		mainMenu(reader, value)
	}
}

func register(reader *bufio.Reader) {
	for {
		pl("Register")
		pl("Enter your username (min 6 characters) or -1 to cancel : ")
		username, err := reader.ReadString('\n')

		username = username[:len(username)-1]

		if err != nil {
			pl("Invalid input", err)
			continue
		}

		if username == "-1" {
			pl("Registration cancelled!\nReturning to login screen...")
			return
		} else if len(username) < 6 {
			pl("Username must be at least 6 characters long!\nPlease try again!")
			continue
		}

		pl("Enter your password (min 8 characters) : ")
		password, err := reader.ReadString('\n')

		password = password[:len(password)-1]

		if err != nil {
			pl("Invalid input", err)
			continue
		}

		if len(password) < 8 {
			pl("Password must be at least 8 characters long!\nPlease try again!")
			continue
		}

		pl("Enter you password again : ")
		password2, err := reader.ReadString('\n')

		password2 = password2[:len(password2)-1]

		if err != nil {
			pl("Invalid input", err)
			continue
		}

		if len(password2) < 8 {
			pl("Password must be at least 8 characters long!\nPlease try again!")
			continue
		} else if password != password2 {
			pl("Passwords do not match!\nPlease try again!")
			continue
		}

		pl("Enter your account type 1 or 2 (1. savings | 2. checkings) : ")
		input, err := reader.ReadString('\n')
		if err != nil {
			pl("Invalid input", err)
			continue
		}

		var accountType string
		input = input[:len(input)-1]

		if input == "1" {
			accountType = "savings"
		} else if input == "2" {
			accountType = "checkings"
		} else {
			pl("Invalid input!\nPlease try again!")
			continue
		}

		reader.ReadString('\n')
		accountNumber := accountNumberGenerator()
		balance := 0.0

		newAccount := Account{
			username:      username,
			password:      password,
			balance:       balance,
			accountType:   accountType,
			accountNumber: accountNumber,
		}
		pl("Account created successfully!")
		pl("Your account number is: ", accountNumber)
		accounts[accountNumber] = newAccount
		accountsByUsername[username] = newAccount
		return
	}
}

func accountNumberGenerator() int {
	return rand.Intn(1000000)
}
