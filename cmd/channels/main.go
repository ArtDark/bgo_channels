package main

import (
	"fmt"
	"github.com/ArtDark/bgo_channels/pkg/card"
	"log"
	"os"
	"runtime/trace"
)

func main() {

	// Trace
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Print(err)
		}
	}()
	err = trace.Start(f)
	if err != nil {
		log.Fatal(err)
	}
	defer trace.Stop()

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
	} // Инициализация карточки пользователя

	err = user.MakeTransactions(50)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("------------------------------------------")
	fmt.Println("Without goroutines")
	fmt.Println(card.SumCategoryTransactions(user.Transactions))
	fmt.Println("------------------------------------------")
	fmt.Println("With mutex")
	fmt.Println(card.SumCategoryTransactionsMutex(user.Transactions, 10))
	fmt.Println("------------------------------------------")
	fmt.Println("With chan")
	fmt.Println(card.SumCategoryTransactionsChan(user.Transactions, 10))
	fmt.Println("------------------------------------------")
	fmt.Println("With mutex without SumCategoryTransactions func")
	fmt.Println(card.SumCategoryTransactionsMutexWithoutFunc(user.Transactions, 10))
	fmt.Println("------------------------------------------")

}
