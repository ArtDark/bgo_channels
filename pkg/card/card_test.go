package card

import (
	"reflect"
	"testing"
)

func TestSumCategoryTransactions(t *testing.T) {
	type args struct {
		transactions []Transaction
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]int64
		wantErr error
	}{
		{
			name: "Valid transactions",
			args: args{
				transactions: []Transaction{
					{Id: "0001", Bill: 100_00, Time: 1606192422, MCC: "5411", Status: "Done"},
					{Id: "0002", Bill: 200_00, Time: 1606192432, MCC: "5812", Status: "Done"},
					{Id: "0003", Bill: 400_00, Time: 1606192442, MCC: "5411", Status: "Done"},
					{Id: "0004", Bill: 300_00, Time: 1606192462, MCC: "5812", Status: "Done"},
				},
			},
			want: map[string]int64{
				"5411": 500_00,
				"5812": 500_00,
			},
			wantErr: nil,
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumCategoryTransactions(tt.args.transactions)
			if err != nil {
				t.Errorf("SumCategoryTransactions() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumCategoryTransactions() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumCategoryTransactionsMutex(t *testing.T) {
	type args struct {
		transactions []Transaction
		goroutines   int
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]int64
		wantErr error
	}{
		{
			name: "Valid transactions",
			args: args{
				transactions: []Transaction{
					{Id: "0001", Bill: 100_00, Time: 1606192422, MCC: "5411", Status: "Done"},
					{Id: "0002", Bill: 200_00, Time: 1606192432, MCC: "5812", Status: "Done"},
					{Id: "0003", Bill: 400_00, Time: 1606192442, MCC: "5411", Status: "Done"},
					{Id: "0004", Bill: 300_00, Time: 1606192462, MCC: "5812", Status: "Done"},
				},
				goroutines: 2,
			},
			want: map[string]int64{
				"5411": 500_00,
				"5812": 500_00,
			},
			wantErr: nil,
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumCategoryTransactionsMutex(tt.args.transactions, tt.args.goroutines)
			if err != tt.wantErr {
				t.Errorf("SumCategoryTransactionsMutex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumCategoryTransactionsMutex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumCategoryTransactionsChan(t *testing.T) {
	type args struct {
		transactions []Transaction
		goroutines   int
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]int64
		wantErr error
	}{
		{
			name: "Valid transactions",
			args: args{
				transactions: []Transaction{
					{Id: "0001", Bill: 100_00, Time: 1606192422, MCC: "5411", Status: "Done"},
					{Id: "0002", Bill: 200_00, Time: 1606192432, MCC: "5812", Status: "Done"},
					{Id: "0003", Bill: 400_00, Time: 1606192442, MCC: "5411", Status: "Done"},
					{Id: "0004", Bill: 300_00, Time: 1606192462, MCC: "5812", Status: "Done"},
				},
				goroutines: 2,
			},
			want: map[string]int64{
				"5411": 500_00,
				"5812": 500_00,
			},
			wantErr: nil,
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumCategoryTransactionsChan(tt.args.transactions, tt.args.goroutines)
			if err != tt.wantErr {
				t.Errorf("SumCategoryTransactionsChan() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumCategoryTransactionsChan() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumCategoryTransactionsMutexWithoutFunc(t *testing.T) {
	type args struct {
		transactions []Transaction
		goroutines   int
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]int64
		wantErr error
	}{
		{
			name: "Valid transactions",
			args: args{
				transactions: []Transaction{
					{Id: "0001", Bill: 100_00, Time: 1606192422, MCC: "5411", Status: "Done"},
					{Id: "0002", Bill: 200_00, Time: 1606192432, MCC: "5812", Status: "Done"},
					{Id: "0003", Bill: 400_00, Time: 1606192442, MCC: "5411", Status: "Done"},
					{Id: "0004", Bill: 300_00, Time: 1606192462, MCC: "5812", Status: "Done"},
				},
				goroutines: 2,
			},
			want: map[string]int64{
				"5411": 500_00,
				"5812": 500_00,
			},
			wantErr: nil,
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SumCategoryTransactionsMutexWithoutFunc(tt.args.transactions, tt.args.goroutines)
			if err != tt.wantErr {
				t.Errorf("SumCategoryTransactionsMutexWithoutFunc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumCategoryTransactionsMutexWithoutFunc() got = %v, want %v", got, tt.want)
			}
		})
	}
}
