// Пакет для работы с транзакциями между картами
package transfer

import (
	"errors"
	"github.com/ArtDark/bgo_card2card/pkg/card"
	"strconv"
	"strings"
)

//Структура сервиса
type Service struct {
	CardSvc       *card.Service
	Commission    float64
	CommissionMin int64
}

//func NewService(cardSvc *card.Service, commission float64, commissionMin int64) *Service {
//	return &Service{CardSvc: cardSvc, Commission: commission, CommissionMin: commissionMin}
//}

var (
	ErrNotEnoughMoney    = errors.New("not enough money")
	ErrInvalidCardNumber = errors.New("wrong card number")
)

// Функция проверки номера карты
func IsValid(n string) bool {
	n = strings.ReplaceAll(n, " ", "") // Удаление пробелов из строки
	if len(n) != 16 {
		return false
	}
	sls := strings.Split(n, "") // Создание слайса из строки
	slsInt := [16]int{}         // Создание слайса типа int

	// Преобразование значение string -> int, запись в слайс int
	for i, j := range sls {
		var err interface{}
		slsInt[i], err = strconv.Atoi(j)
		if err != nil {
			return false
		}
	}
	// Операция над каждым нечетным числом с последующим изменением в слайсе slsInt
	for i := 0; i < len(slsInt); i += 2 {
		num := slsInt[i] * 2

		if num > 9 {
			num -= 9
		}

		slsInt[i] = num
	}

	sum := 0 // Контрольная сумма

	// Сумма всех чисел в слайсе
	for _, i := range slsInt {
		sum += i
	}

	// Проверка на кратность 10
	if sum%10 == 0 {
		return true
	}

	return false
}

// Функция перевода с карты на карту
func (s *Service) Card2Card(from, to string, amount int) (int, error) {

	if !IsValid(from) || !IsValid(to) {
		return amount, ErrInvalidCardNumber
	}

	commission := float64(amount) * s.Commission / 100.0 //Расчет комиссии
	total := amount + int(commission)                    // Расчет суммы перевода с комиссией

	toCard, err := s.CardSvc.Card(to)
	if err != nil {
		toCard.Balance += amount
		return total, err

	}
	fromCard, err := s.CardSvc.Card(from) // Поиск карты отправителя

	// Поиск карты получателя
	if fromCard == nil && toCard == nil {
		return total, nil
	}

	if toCard != nil && fromCard.Balance >= amount {
		fromCard.Balance -= int(float64(amount) + commission)
		return total, nil
	}

	if amount > fromCard.Balance { // Если баланс меньше суммы
		return amount, ErrNotEnoughMoney
	}

	fromCard.Balance -= total
	toCard.Balance += amount

	return total, nil

}
