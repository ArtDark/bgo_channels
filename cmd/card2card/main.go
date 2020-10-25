package main

import (
	"fmt"

	"github.com/ArtDark/bgo_channels/pkg/card"
	"github.com/ArtDark/bgo_channels/pkg/transfer"
)

func main() {

	svc := card.New("YourBank")

	visa := svc.CardIssue(
		"0001",
		"Artem",
		"Balusov",
		"Visa",
		12345_67,
		"RUR",
		"5106 2142 2342 4322",
	)

	master := svc.CardIssue(
		"0002",
		"Ivan",
		"Ivanov",
		"MasterCard",
		98765_43,
		"RUR",
		"5106 2142 4322 2342",
	)

	fmt.Println(svc)
	fmt.Println(visa)
	fmt.Println(master)

	s := transfer.IsValid("4  5  6  1     2  6  1  2     1  2  3  4     5  4  6  7")
	//
	//transferSrv := transfer.NewService(svc, 0.5, 10)
	//
	//total, status := transferSrv.Card2Card("1233_2342_2342_4322", "3242_3242_4322_2342", 5000_00)
	//
	//fmt.Println(total, status)
	//fmt.Println(visa)
	//fmt.Println(master)
	fmt.Println(s)

}
