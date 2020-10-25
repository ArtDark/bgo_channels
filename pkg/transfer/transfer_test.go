package transfer

import (
	"github.com/ArtDark/bgo_card2card/pkg/card"
	"testing"
)

func TestService_Card2Card(t *testing.T) {
	type fields struct {
		CardSvc       *card.Service
		Commission    float64
		CommissionMin int64
	}
	type args struct {
		from   string
		to     string
		amount int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTotal int
		wantOk    error
	}{
		{
			name: "Карта своего банка -> Карта своего банка (денег достаточно)",
			fields: fields{
				CardSvc: &card.Service{
					BankName: "YourBank",
					Cards: []*card.Card{
						{
							Id: "0001",
							Owner: card.Owner{
								FirstName: "Artem",
								LastName:  "Balusov",
							},
							Issuer:   "Visa",
							Balance:  43534_34,
							Currency: "RUR",
							Number:   "5106 2158 3920 4837",
							Icon:     "http://...",
						},
						{
							Id: "0002",
							Owner: card.Owner{
								FirstName: "Ivan",
								LastName:  "Ivanov",
							},
							Issuer:   "MasterCard",
							Balance:  543534_23,
							Currency: "RUR",
							Number:   "5106 2163 6456 3456",
							Icon:     "http://...",
						},
					},
				},
				Commission:    0.5,
				CommissionMin: 10,
			},
			args: args{
				from:   "5106 2158 3920 4837",
				to:     "5106 2163 6456 3456",
				amount: 1000_00,
			},
			wantTotal: 1005_00,
			wantOk:    nil,
		},
		{
			name: "Карта своего банка -> Карта своего банка (денег недостаточно)",
			fields: fields{
				CardSvc: &card.Service{
					BankName: "YourBank",
					Cards: []*card.Card{
						{
							Id: "0001",
							Owner: card.Owner{
								FirstName: "Artem",
								LastName:  "Balusov",
							},
							Issuer:   "Visa",
							Balance:  34_34,
							Currency: "RUR",
							Number:   "5106 2158 3920 4837",
							Icon:     "http://...",
						},
						{
							Id: "0002",
							Owner: card.Owner{
								FirstName: "Ivan",
								LastName:  "Ivanov",
							},
							Issuer:   "MasterCard",
							Balance:  543534_23,
							Currency: "RUR",
							Number:   "5106 2163 6456 3456",
							Icon:     "http://...",
						},
					},
				},
				Commission:    0.5,
				CommissionMin: 10,
			},
			args: args{
				from:   "5106 2158 3920 4837",
				to:     "5106 2163 6456 3456",
				amount: 1000_00,
			},
			wantTotal: 1000_00,
			wantOk:    ErrNotEnoughMoney,
		},
		{
			name: "Карта своего банка -> Карта чужого банка (денег достаточно)",
			fields: fields{
				CardSvc: &card.Service{
					BankName: "YourBank",
					Cards: []*card.Card{
						{
							Id: "0001",
							Owner: card.Owner{
								FirstName: "Artem",
								LastName:  "Balusov",
							},
							Issuer:   "Visa",
							Balance:  43534_34,
							Currency: "RUR",
							Number:   "5106 2158 3920 4837",
							Icon:     "http://...",
						},
						{
							Id: "0002",
							Owner: card.Owner{
								FirstName: "Ivan",
								LastName:  "Ivanov",
							},
							Issuer:   "MasterCard",
							Balance:  543534_23,
							Currency: "RUR",
							Number:   "5106 2163 6456 3456",
							Icon:     "http://...",
						},
					},
				},
				Commission:    0.5,
				CommissionMin: 10,
			},
			args: args{
				from:   "5106 2158 3920 4837",
				to:     "4578 8942 5433 4329",
				amount: 1000_00,
			},
			wantTotal: 1005_00,
			wantOk:    nil,
		},
		{
			name: "Карта своего банка -> Карта чужого банка (денег недостаточно)",
			fields: fields{
				CardSvc: &card.Service{
					BankName: "YourBank",
					Cards: []*card.Card{
						{
							Id: "0001",
							Owner: card.Owner{
								FirstName: "Artem",
								LastName:  "Balusov",
							},
							Issuer:   "Visa",
							Balance:  34_34,
							Currency: "RUR",
							Number:   "5106 2158 3920 4837",
							Icon:     "http://...",
						},
						{
							Id: "0002",
							Owner: card.Owner{
								FirstName: "Ivan",
								LastName:  "Ivanov",
							},
							Issuer:   "MasterCard",
							Balance:  543534_23,
							Currency: "RUR",
							Number:   "5106 2163 6456 3456",
							Icon:     "http://...",
						},
					},
				},
				Commission:    0.5,
				CommissionMin: 10,
			},
			args: args{
				from:   "5106 2158 3920 4837",
				to:     "4578 8942 5433 4329",
				amount: 1000_00,
			},
			wantTotal: 1000_00,
			wantOk:    ErrNotEnoughMoney,
		},
		{
			name: "Карта чужого банка -> Карта своего банка",
			fields: fields{
				CardSvc: &card.Service{
					BankName: "YourBank",
					Cards: []*card.Card{
						{
							Id: "0001",
							Owner: card.Owner{
								FirstName: "Artem",
								LastName:  "Balusov",
							},
							Issuer:   "Visa",
							Balance:  43534_34,
							Currency: "RUR",
							Number:   "5106 2158 3920 4837",
							Icon:     "http://...",
						},
						{
							Id: "0002",
							Owner: card.Owner{
								FirstName: "Ivan",
								LastName:  "Ivanov",
							},
							Issuer:   "MasterCard",
							Balance:  543534_23,
							Currency: "RUR",
							Number:   "5106 2163 6456 3456",
							Icon:     "http://...",
						},
					},
				},
				Commission:    0.5,
				CommissionMin: 10,
			},
			args: args{
				from:   "4578 8942 5433 4329",
				to:     "5106 2163 6456 3456",
				amount: 1000_00,
			},
			wantTotal: 1005_00,
			wantOk:    nil,
		},
		{
			name: "Карта чужого банка -> Карта чужого банка",
			fields: fields{
				CardSvc: &card.Service{
					BankName: "YourBank",
					Cards: []*card.Card{
						{
							Id: "0001",
							Owner: card.Owner{
								FirstName: "Artem",
								LastName:  "Balusov",
							},
							Issuer:   "Visa",
							Balance:  43534_34,
							Currency: "RUR",
							Number:   "5106 2158 3920 4837",
							Icon:     "http://...",
						},
						{
							Id: "0002",
							Owner: card.Owner{
								FirstName: "Ivan",
								LastName:  "Ivanov",
							},
							Issuer:   "MasterCard",
							Balance:  543534_23,
							Currency: "RUR",
							Number:   "5106 2163 6456 3456",
							Icon:     "http://...",
						},
					},
				},
				Commission:    0.5,
				CommissionMin: 10,
			},
			args: args{
				from:   "4578 8942 5433 4329",
				to:     "4534 5963 6456 3456",
				amount: 1000_00,
			},
			wantTotal: 1005_00,
			wantOk:    nil,
		},
		{
			name: "Неправильная карта своего банка -> Карта чужого банка",
			fields: fields{
				CardSvc: &card.Service{
					BankName: "YourBank",
					Cards: []*card.Card{
						{
							Id: "0001",
							Owner: card.Owner{
								FirstName: "Artem",
								LastName:  "Balusov",
							},
							Issuer:   "Visa",
							Balance:  43534_34,
							Currency: "RUR",
							Number:   "5106 2158 3920 4837",
							Icon:     "http://...",
						},
						{
							Id: "0002",
							Owner: card.Owner{
								FirstName: "Ivan",
								LastName:  "Ivanov",
							},
							Issuer:   "MasterCard",
							Balance:  543534_23,
							Currency: "RUR",
							Number:   "5106 2163 6456 3456",
							Icon:     "http://...",
						},
					},
				},
				Commission:    0.5,
				CommissionMin: 10,
			},
			args: args{
				from:   "5106 2742 5433 4321",
				to:     "4578 8942 5433 4329",
				amount: 1000_00,
			},
			wantTotal: 1000_00,
			wantOk:    ErrInvalidCardNumber,
		},
		{
			name: "Карта чужого банка -> Неправильная карта своего банка",
			fields: fields{
				CardSvc: &card.Service{
					BankName: "YourBank",
					Cards: []*card.Card{
						{
							Id: "0001",
							Owner: card.Owner{
								FirstName: "Artem",
								LastName:  "Balusov",
							},
							Issuer:   "Visa",
							Balance:  43534_34,
							Currency: "RUR",
							Number:   "5106 2158 3920 4837",
							Icon:     "http://...",
						},
						{
							Id: "0002",
							Owner: card.Owner{
								FirstName: "Ivan",
								LastName:  "Ivanov",
							},
							Issuer:   "MasterCard",
							Balance:  543534_23,
							Currency: "RUR",
							Number:   "5106 2163 6456 3456",
							Icon:     "http://...",
						},
					},
				},
				Commission:    0.5,
				CommissionMin: 10,
			},
			args: args{
				from:   "4578 8942 5433 4329",
				to:     "5106 2163 6856 3456",
				amount: 1000_00,
			},
			wantTotal: 1000_00,
			wantOk:    ErrInvalidCardNumber,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CardSvc:       tt.fields.CardSvc,
				Commission:    tt.fields.Commission,
				CommissionMin: tt.fields.CommissionMin,
			}
			gotTotal, gotOk := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
			if gotTotal != tt.wantTotal {
				t.Errorf("Card2Card() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Card2Card() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestIsValid(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Неправильные символы",
			args: args{"1233 2132 2131 213q"},
			want: false,
		},
		{
			name: "Правильный номер карты",
			args: args{"4561 2612 1234 5467"},
			want: true,
		},
		{
			name: "Неправильный номер карты",
			args: args{"4561 2662 1234 5467"},
			want: false,
		},
		{
			name: "Неправильная длина номера карты",
			args: args{"4561 2612 12374 5467"},
			want: false,
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.n); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
