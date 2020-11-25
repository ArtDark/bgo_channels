package main

import (
	"fmt"
	"github.com/ArtDark/bgo_channels/pkg/card"
)

func main() {
	// Инициализация карточки пользователя
	user := card.Card{
		Id: 1,
		Owner: card.Owner{
			FirstName: "Ivan",
			LastName:  "Petrov",
		},
		Issuer:       "Master Card",
		Balance:      48234_63,
		Currency:     "RUB",
		Number:       "5106212365738734",
		Icon:         "https://www.mastercard.ru/content/dam/public/enterprise/resources/images/icons/favicon.ico",
		Transactions: []card.Transaction{},
	}

	err := user.MakeTransactions(50)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("------------------------------------------")
	fmt.Println("Without goroutines")
	fmt.Println(card.SumCategoryTransactions(user.Transactions))
	fmt.Println("------------------------------------------")
	fmt.Println("With mutex")
	fmt.Println(card.SumCategoryTransactionsMutex(user.Transactions, 3))
	fmt.Println("------------------------------------------")
	fmt.Println("With chan")
	fmt.Println(card.SumCategoryTransactionsChan(user.Transactions, 3))
	fmt.Println("------------------------------------------")
	fmt.Println("With mutex without SumCategoryTransactions func")
	fmt.Println(card.SumCategoryTransactionsMutexWithoutFunc(user.Transactions, 3))
	fmt.Println("------------------------------------------")

}
