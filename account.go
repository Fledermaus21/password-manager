package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

type account struct {
	login    string
	password string
	url      string
}

func newAccount(login, password, urlString string) (*account, error) {
	//Validation
	if login == "" {
		return nil, errors.New("No login")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Invalid URL")
	}

	newAcc := &account{
		login:    login,
		password: password,
		url:      urlString,
	}

	if password == "" {
		newAcc.generatePassword(10)
	}

	return newAcc, nil
}

func (acc account) outputData() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	symbols := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	result := []rune{}
	for i := 0; i < n; i++ {
		result = append(result, symbols[rand.IntN(62)])
	}
	acc.password = string(result)
}
