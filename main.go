package main

import (
	"bufio"
	"fmt"
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
	case 1:
		login()
	case 2:
		register()
	case 3:
		pl("Exiting...")
		return
	default:
		pl("Invalid input")
	}
}

func login() {
	reader := bufio.NewReader(os.Stdin)
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

}

func register() {
	while true {
		reader := bufio.NewReader(os.Stdin)
		pl("Register")
		pl("Enter your username: ")
		username, err := reader.ReadString('\n')
		if err != nil {
			pl("Invalid input", err)
			break
		}
	}
}
