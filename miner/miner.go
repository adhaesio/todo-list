package miner

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Miner(ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- int, n int,
	power int) {

	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Я шахтер номер:", n, "Рабочий день окончен!")
			return
		default:
			fmt.Println("Я шахтер номер:", n, "Начал добывать уголь!")
			time.Sleep(1 * time.Second)
			fmt.Println("Я шахтёр номер:", n, "Добыл угля", power)
			transferPoint <- power
			fmt.Println("Я шахтёр номер:", n, "Передал уголь", power)

		}

	}

}

func MinerPool(ctx context.Context, minerCount int) <-chan int {
	coaltransferpoint := make(chan int)
	wg := &sync.WaitGroup{}
	for i := 1; i <= minerCount; i++ {
		wg.Add(1)
		go Miner(ctx, wg, coaltransferpoint, i, i*10)

	}
	go func() {
		wg.Wait()
		close(coaltransferpoint)
	}()
	return coaltransferpoint

}
