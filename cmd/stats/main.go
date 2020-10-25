package main

import (
	"fmt"
	"github.com/ArtDark/bgo_goroutines/pkg/stats"
	"log"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

func main() {
	//runtime.GOMAXPROCS(4)

	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Print(err)
		}
	}()
	err = trace.Start(f)
	if err != nil {
		log.Fatal(err)
	}
	defer trace.Stop()

	const users = 10_000_000
	const transactionsPerUser = 100
	const transactionAmount = 1_00
	transactions := make([]int64, users*transactionsPerUser)
	for index := range transactions {
		// для простоты храним только суммы
		// и считаем, что каждая транзакция = 1 руб.
		transactions[index] = transactionAmount
	}

	start := time.Date(2015, 4, 1, 0, 0, 0, 0, time.Local)
	finish := time.Date(2016, 4, 1, 0, 0, 0, 0, time.Local)

	months := make([]string, 0)
	next := start
	for next.Before(finish) {
		months = append(months, fmt.Sprintf("%d-%d", next.Year(), next.Month()))
		next = next.AddDate(0, 1, 0)
	}

	monthsMap := make(map[string][]int64)
	partsCount := len(months)
	partSize := len(transactions) / partsCount

	for i := 0; i < len(months); i++ {
		monthsMap[months[i]] = transactions[i*partSize : (i+1)*partSize]
	}

	wg := sync.WaitGroup{}

	for m, i := range monthsMap {
		wg.Add(1)
		go func(month string, amounts []int64) {
			sum := stats.Sum(amounts)
			fmt.Printf("Sum month %s - %d\n", month, sum)
			wg.Done()
		}(m, i)
	}

	wg.Wait()

}
