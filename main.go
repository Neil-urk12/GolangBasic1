package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

var pl = fmt.Println

type Account struct {
	username      string
	password      string
	balance       float64
	accountType   string
	accountNumber int
}

var accounts []Account

func main() {
	reader := bufio.NewReader(os.Stdin)
	pl("==Welcome to Golangking System!==")
	loginScreen(reader)
}

func loginScreen(reader *bufio.Reader) {
	pl("1. Login  2. Register  3. Exit")
	pl("Enter you choice: ")
	choice, err := reader.ReadByte()
	if err != nil {
		pl("Invalid input")
		return
	}
	switch choice {
	case '1':
		login(reader)
	case '2':
		register(reader)
	case '3':
		pl("Exiting...")
		return
	default:
		pl("Invalid input")
	}
}

func login(reader *bufio.Reader) {
	pl("Login")
	pl("Enter your username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		pl("Invalid input", err)
		return
	}
	pl("Enter your password: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		pl("Invalid input", err)
		return
	}
	pl(username, password)
}

func register(reader *bufio.Reader) {
	for {
		pl("Register")

		pl("Enter your username (min 8 characters): ")
		username, err := reader.ReadString('\n')
		if err != nil {
			pl("Invalid input", err)
			continue
		} else if len(username) < 8 {
			pl("Username must be at least 8 characters long!\nPlease try again!")
			continue
		}

		pl("Enter your password (min 8 characters) : ")
		password, err := reader.ReadString('\n')
		if err != nil {
			pl("Invalid input", err)
			continue
		} else if len(password) < 8 {
			pl("Password must be at least 8 characters long!\nPlease try again!")
			continue
		}

		pl("Enter you password again : ")
		password2, err := reader.ReadString('\n')
		if err != nil {
			pl("Invalid input", err)
			continue
		} else if len(password2) < 8 {
			pl("Password must be at least 8 characters long!\nPlease try again!")
			continue
		} else if password != password2 {
			pl("Passwords do not match!\nPlease try again!")
			continue
		}

		pl("Enter your account type (savings/checkings) : ")
		accountType, err := reader.ReadString('\n')
		if err != nil {
			pl("Invalid input", err)
			continue
		}
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
		accounts = append(accounts, newAccount)
	}
}

func accountNumberGenerator() int {
	return rand.Intn(1000000)
}
