package main

import (
	"github.com/ArtDark/bgo_channels/pkg/card"
	"strconv"
	"time"
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

	// Генератор транзакци TODO: можно переделать в функцию для удобства
	transactionCounts := 10 // Количество генереци 2-х транзакций с разными MCC

	for i := 0; i < transactionCounts; i++ {
		user.AddTransaction(card.Transaction{
			Id: strconv.Itoa((i + 1) + i),

			Bill: int64(100_00 + i*10),

			Time:   time.Date(2020, 9, 10, 12+i, 23+i, 21+i, 0, time.UTC).Unix(),
			MCC:    "5411",
			Status: "Done",
		})
		user.AddTransaction(card.Transaction{
			Id: strconv.Itoa((i + 2) + i),

			Bill: int64(102_00 + i*10),

			Time:   time.Date(2020, 9, 10, 14+i, 15+i, 21+i, 0, time.UTC).Unix(),
			MCC:    "5812",
			Status: "Done",
		})

	}

	// Перечисление списка транзакций
	//for t, _ := range user.Transactions {
	//	fmt.Println(user.Transactions[t])
	//}

	//fmt.Println(card.SumCategoryTransactions(user.Transactions))
	card.SumCategoryTransactionsMutex(user.Transactions, 3)



}

func SumChan(user.Transactions, result chan <- map[string]int64)  {
	go func() {
		result <- card.SumCategoryTransactionChan(transactions)
	}()
}
