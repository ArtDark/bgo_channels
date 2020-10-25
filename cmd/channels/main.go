package main

import (
	"fmt"
	"github.com/ArtDark/bgo_channels/pkg/card"
	"time"
)

func main() {

	user := card.Card{
		Id: "0001",
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

	for i := 0; i < 4; i++ {
		user.AddTransaction(card.Transaction{
			Id:     string(48 + (i + 1) + i),
			Bill:   int64(345_00 + i*100),
			Time:   time.Date(2020, 9, 10, 12+i, 23+i, 21+i, 0, time.UTC).Unix(),
			MCC:    "5411",
			Status: "Done",
		})
		user.AddTransaction(card.Transaction{
			Id:     string(48 + (i + 2) + i),
			Bill:   int64(345_00 + (i+2)*100),
			Time:   time.Date(2020, 9, 10, 14+i, 15+i, 21+i, 0, time.UTC).Unix(),
			MCC:    "5812",
			Status: "Done",
		})

	}

	fmt.Println(user)

}
