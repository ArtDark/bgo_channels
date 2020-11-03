//Package card
package card

import (
	"errors"
	"fmt"
	"strings"
	"sync"
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

func (card *Card) AddTransaction(transaction Transaction) {
	card.Transactions = append(card.Transactions, transaction)

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

// TODO: Обычная функция, которая принимает на вход слайс транзакций и id владельца - возвращает map с категориям и тратами по ним (сортировать они ничего не должна)

func SumCategoryTransactions(transactions []Transaction) (map[string]int64, error) {

	if transactions == nil {
		return nil, ErrNoTransactions
	}

	m := make(map[string]int64)

	for j, _ := range transactions {
		m[transactions[j].MCC] += transactions[j].Bill
	}

	return m, nil

}

// TODO: Функция с mutex'ом, который защищает любые операции с map, соответственно, её задача: разделить слайс транзакций на несколько кусков и в отдельных горутинах посчитать map'ы по кускам, после чего собрать всё в один большой map. Важно: эта функция внутри себя должна вызывать функцию из п.1

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
			mapSum, _ := SumCategoryTransactions(part)
			mu.Lock()
			for key, i := range mapSum {
				m[key] += i
				fmt.Println(m)

			}
			mu.Unlock()
			wg.Done()
		}()

	}
	wg.Wait()

	return m, nil

}

// TODO: Функция с каналами, соответственно, её задача: разделить слайс транзакций на несколько кусков и в отдельных горутинах посчитать map'ы по кускам, после чего собрать всё в один большой map (передавайте рассчитанные куски по каналу). Важно: эта функция внутри себя должна вызывать функцию из п.1
func SumCategoryTransactionsChan(transactions []Transaction, goroutines int) (map[string]int64, error) {

	if transactions == nil {
		return nil, ErrNoTransactions
	}

	m := make(map[string]int64)

	partSize := len(transactions) / goroutines

	for i := 0; i < goroutines; i++ {
		part := transactions[i*partSize : (i+1)*partSize]
		go func() {
			mapSum, _ := SumCategoryTransactions(part)
			for key, i := range mapSum {
				m[key] += i
				fmt.Println(m)

			}

		}()

	}

	return m, nil

}

// TODO: Функция с mutex'ом, который защищает любые операции с map, соответственно, её задача: разделить слайс транзакций на несколько кусков и в отдельных горутинах посчитать, но теперь горутины напрямую пишут в общий map с результатами. Важно: эта функция внутри себя не должна вызывать функцию из п.1

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

var (
	ErrCardNotFound   = errors.New("card not found")
	ErrNoTransactions = errors.New("no user transactions")
)

const prefix = "5106 21" //Первые 6 цифр нашего банка

// Метод поиска банковской карты по номеру платежной системы
func (s *Service) Card(number string) (*Card, error) {

	for _, с := range s.Cards {
		if strings.HasPrefix(с.Number, prefix) == true {
			return с, nil
		}
	}
	return nil, ErrCardNotFound
}
