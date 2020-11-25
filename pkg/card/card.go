//Package card
package card

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	ErrCardNotFound   = errors.New("card not found")
	ErrNoTransactions = errors.New("no user transactions")
)

// Описание банковской карты"
type Card struct {
	Id           cardId
	Owner               // Владелец карты
	Issuer       string // Платежная истема
	Balance      int    // Баланс карты
	Currency     string // Валюта
	Number       string // Номер карты в платежной системе
	Icon         string // Иконка платежной системы
	Transactions []Transaction
}

// Идентификат банковской карты
type cardId int64

// Инициалы владельца банковской карты
type Owner struct {
	FirstName string // Имя владельца карты
	LastName  string // Фамилия владельца карты
}

type Transaction struct {
	Id     string
	Bill   int64
	Time   int64
	MCC    string
	Status string
}

// Метод добавления транзакции
func (card *Card) AddTransaction(transaction Transaction) {
	card.Transactions = append(card.Transactions, transaction)

}

// Метод геренерации 2х транзакций с разными MCC
func (card *Card) MakeTransactions(count int) error {

	if card == nil {
		return ErrCardNotFound
	}

	if count <= 0 {
		log.Println("count must be > 0")
		return nil
	}

	for i := 0; i < count; i++ {
		card.AddTransaction(Transaction{
			Id: strconv.Itoa((i + 1) + i),

			Bill: int64(100_00 + i*10),

			Time:   time.Date(2020, 9, 10, 12+i, 23+i, 21+i, 0, time.UTC).Unix(),
			MCC:    "5411",
			Status: "Done",
		})
		card.AddTransaction(Transaction{
			Id: strconv.Itoa((i + 2) + i),

			Bill: int64(102_00 + i*10),

			Time:   time.Date(2020, 9, 10, 14+i, 15+i, 21+i, 0, time.UTC).Unix(),
			MCC:    "5812",
			Status: "Done",
		})

	}

	return nil

}

// Функция расчета суммы по категории
func SumByMCC(transactions []Transaction, mcc []string) int64 {
	var mmcSum int64

	for _, code := range mcc {
		for _, t := range transactions {
			if code == t.MCC {
				mmcSum += t.Bill
			}
		}
	}

	return mmcSum

}

// Функция преобразования кода в название категории
func TranslateMCC(code string) string {
	// представим, что mcc читается из файла (научимся позже)
	mcc := map[string]string{
		"5411": "Супермаркеты",
		"5812": "Рестораны",
	}

	const errCategoryUndef = "Категория не указана"

	if value, ok := mcc[code]; ok {
		return value
	}

	return errCategoryUndef

}

// Функция сложения сумм транзакций по категориям
func SumCategoryTransactions(transactions []Transaction) (map[string]int64, error) {

	if transactions == nil {
		return nil, ErrNoTransactions
	}

	m := make(map[string]int64)

	for i := range transactions {
		m[transactions[i].MCC] += transactions[i].Bill
	}

	return m, nil

}

// Функция сложения сумм транзакций по категориям с использованием goroutines и mutex
func SumCategoryTransactionsMutex(transactions []Transaction, goroutines int) (map[string]int64, error) {
	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	mu := sync.Mutex{}

	if transactions == nil {
		return nil, ErrNoTransactions
	}

	m := make(map[string]int64)

	partSize := len(transactions) / goroutines

	for i := 0; i < goroutines; i++ {
		part := transactions[i*partSize : (i+1)*partSize]
		go func() {
			mapSum, err := SumCategoryTransactions(part)
			if err != nil {
				fmt.Println(err)
			}
			mu.Lock()
			for key, i := range mapSum {
				m[key] += i

			}
			mu.Unlock()
			wg.Done()
		}()

	}
	wg.Wait()

	return m, nil

}

// Функция сложения сум транзакций по категориям с использованием goroutines и каналов
func SumCategoryTransactionsChan(transactions []Transaction, goroutines int) (map[string]int64, error) {

	if transactions == nil {
		return nil, ErrNoTransactions
	}

	result := make(map[string]int64)
	ch := make(chan map[string]int64)
	partSize := len(transactions) / goroutines

	for i := 0; i < goroutines; i++ {
		part := transactions[i*partSize : (i+1)*partSize]
		go func(ch chan<- map[string]int64) {
			s, err := SumCategoryTransactions(part)
			if err != nil {
				log.Printf("failed to sum: %s\n", err)
			}
			ch <- s
		}(ch)
	}

	fin := 0

	for sum := range ch {
		for k, v := range sum {
			result[k] += v

		}
		fin++
		if fin == goroutines {
			close(ch)
			break
		}
	}

	return result, nil

}

// Функция с mutex'ом, который защищает любые операции с map, соответственно, её задача: разделить слайс транзакций на несколько кусков и в отдельных горутинах посчитать, но теперь горутины напрямую пишут в общий map с результатами. Важно: эта функция внутри себя не должна вызывать функцию из п.1
func SumCategoryTransactionsMutexWithoutFunc(transactions []Transaction, goroutines int) (map[string]int64, error) {
	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	mu := sync.Mutex{}

	if transactions == nil {
		return nil, ErrNoTransactions
	}

	mapSum := make(map[string]int64)

	partSize := len(transactions) / goroutines

	for i := 0; i < goroutines; i++ {
		part := transactions[i*partSize : (i+1)*partSize]
		go func() {

			for i := range part {
				mu.Lock()
				mapSum[part[i].MCC] += part[i].Bill
				mu.Unlock()
			}
			wg.Done()

		}()

	}
	wg.Wait()

	return mapSum, nil

}

// Сервис банка
type Service struct {
	BankName string
	Cards    []*Card
}

// Конструктор сервиса
func New(bankName string) *Service {
	return &Service{BankName: bankName}
}

// Метод создания экземпляра банковской карты
func (s *Service) CardIssue(
	id cardId,
	fistName,
	lastName,
	issuer string,
	balance int,
	currency string,
	number string,
) *Card {
	var card = &Card{
		Id: id,
		Owner: Owner{
			FirstName: fistName,
			LastName:  lastName,
		},
		Issuer:   issuer,
		Balance:  balance,
		Currency: currency,
		Number:   number,
		Icon:     "https://.../logo.png",
	}
	s.Cards = append(s.Cards, card)
	return card
}

const prefix = "5106 21" //Первые 6 цифр нашего банка

// Метод поиска банковской карты по номеру платежной системы
func (s *Service) Card() (*Card, error) {

	for _, c := range s.Cards {
		if strings.HasPrefix(c.Number, prefix) == true {
			return c, nil
		}
	}
	return nil, ErrCardNotFound
}
