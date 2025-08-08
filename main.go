package main

import (
	"fmt"
	"password-manager/account"
	"password-manager/files"
)

func main() {
	vault := account.NewVault()
Menu:
	for {
		fmt.Println("Выберите опцию")
		fmt.Println("1. Создать аккаунт")
		fmt.Println("2. Удалить аккаунт")
		fmt.Println("3. Найти аккаунт по URL")
		fmt.Println("4. Выйти")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			createAccount(vault)
		case 2:
			deleteAccount(vault)
		case 3:
			findAccount(vault)
		default:
			break Menu
		}
	}
}

func deleteAccount(vault *account.Vault) {
	login := promptData("Введите логин: ")
	url := promptData("Введите URL: ")

	result := vault.DeleteAccount(login, url)
	if result {
		fmt.Println("Успешно удалено")
	} else {
		fmt.Println("Не найдено")
	}
}

func findAccount(vault *account.Vault) {
	url := promptData("Введите URL: ")
	accounts := vault.FindAccByURL(url)
	if len(accounts) == 0 {
		fmt.Println("Аккаунтов не найдено")
	}
	for _, account := range accounts {
		account.OutputData()
	}
}

func createAccount(vault *account.Vault) {
	login := promptData("Введите логин: ")
	password := promptData("Введите пароль: ")
	url1 := promptData("Введите URL: ")

	account1, err := account.NewAccount(login, password, url1)
	if err != nil {
		fmt.Println(err)
		return
	}
	account1.OutputData()

	vault.AddAccount(*account1)
	data, err := vault.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать в JSON")
		return
	}
	files.WriteFile(data, "data.json")
}

func promptData(prompt string) string {
	var result string
	fmt.Println(prompt)
	fmt.Scanln(&result)
	return result
}
