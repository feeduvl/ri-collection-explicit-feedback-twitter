package main

import "fmt"

type ResponseMessage struct {
	AccountExists bool   `json:"account_exists"`
	Message       string `json:"message"`
}

func (r ResponseMessage) Create(accountExists bool, accountName string) ResponseMessage {
	r.AccountExists = accountExists
	if accountExists {
		r.Message = fmt.Sprintf("Account %s found on Twitter", accountName)
	} else {
		r.Message = fmt.Sprintf("Account %s does not exist", accountName)
	}

	return r
}
