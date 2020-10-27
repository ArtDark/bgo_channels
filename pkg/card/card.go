//Package card
package card

import (
	"errors"
	"strings"
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

func (c *Card) AddTransaction(transaction Transaction) {
	c.Transactions = append(c.Transactions, transaction)
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

var ErrCardNotFound = errors.New("card not found")

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
