package main

import (
	"fmt"
)

func main() {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url1 := promptData("Введите URL: ")

	account1, err := newAccount(login, password, url1)
	if err != nil {
		fmt.Println(err)
		return
	}
	account1.generatePassword(10)
	fmt.Println(account1.password)
	account1.outputData()
}

func promptData(prompt string) string {
	var result string
	fmt.Println(prompt)
	fmt.Scanln(&result)
	return result
}
