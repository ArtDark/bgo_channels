package main

import (
	"fmt"
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
	transactionCounts := 100000 // Количество генереци 2-х транзакций с разными MCC

	for i := 0; i < transactionCounts; i++ {
		user.AddTransaction(card.Transaction{
			Id:     strconv.Itoa((i + 1) + i),
			Bill:   int64(300_00 + i*10),
			Time:   time.Date(2020, 9, 10, 12+i, 23+i, 21+i, 0, time.UTC).Unix(),
			MCC:    "5411",
			Status: "Done",
		})
		user.AddTransaction(card.Transaction{
			Id:     strconv.Itoa((i + 2) + i),
			Bill:   int64(400_00 + i*10),
			Time:   time.Date(2020, 9, 10, 14+i, 15+i, 21+i, 0, time.UTC).Unix(),
			MCC:    "5812",
			Status: "Done",
		})

	}

	// Перечисление списка транзакций
	for t, _ := range user.Transactions {
		fmt.Println(user.Transactions[t])
	}

	// TODO: Обычная функция, которая принимает на вход слайс транзакций и id владельца - возвращает map с категориям и тратами по ним (сортировать они ничего не должна)

	// TODO: Функция с mutex'ом, который защищает любые операции с map, соответственно, её задача: разделить слайс транзакций на несколько кусков и в отдельных горутинах посчитать map'ы по кускам, после чего собрать всё в один большой map. Важно: эта функция внутри себя должна вызывать функцию из п.1

	// TODO: Функция с каналами, соответственно, её задача: разделить слайс транзакций на несколько кусков и в отдельных горутинах посчитать map'ы по кускам, после чего собрать всё в один большой map (передавайте рассчитанные куски по каналу). Важно: эта функция внутри себя должна вызывать функцию из п.1

	// TODO: Функция с mutex'ом, который защищает любые операции с map, соответственно, её задача: разделить слайс транзакций на несколько кусков и в отдельных горутинах посчитать, но теперь горутины напрямую пишут в общий map с результатами. Важно: эта функция внутри себя не должна вызывать функцию из п.1

}
