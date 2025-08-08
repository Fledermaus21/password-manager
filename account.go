package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewAccount(login, password, urlString string) (*Account, error) {
	//Validation
	if login == "" {
		return nil, errors.New("No login")
	}

	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("Invalid URL")
	}

	newAcc := &Account{
		Login:     login,
		Password:  password,
		Url:       urlString,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if password == "" {
		newAcc.GeneratePassword(10)
	}

	return newAcc, nil
}

func (acc Account) OutputData() {
	fmt.Println(acc.Login, acc.Password, acc.Url)
}

func (acc *Account) GeneratePassword(n int) {
	symbols := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	result := []rune{}
	for i := 0; i < n; i++ {
		result = append(result, symbols[rand.IntN(62)])
	}
	acc.Password = string(result)
}
