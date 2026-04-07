package postman

import (
	"context"
	"fmt"

	"sync"
	"time"
)

func Postman(ctx context.Context, wg *sync.WaitGroup, transferPointn chan<- string, n int, mail string) {

	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Я почтальон номер:", n, "Я заебался работать я заканчиваю тут")
			return
		default:
			fmt.Println("Я почтальон номер:", n, "Взял письмо")
			time.Sleep(1 * time.Second)
			fmt.Println("Я почтальон номер:", n, "Донёс письмо до почты:", mail)
			transferPointn <- mail
			fmt.Println("Я почтальон номер:", n, "Передал письмо:", mail)
		}

	}

}

func PostmanPool(ctx context.Context, postmanCount int) <-chan string {
	mailTransferPoint := make(chan string)
	wg := &sync.WaitGroup{}
	for i := 1; i <= postmanCount; i++ {
		wg.Add(1)
		go Postman(ctx, wg, mailTransferPoint, i, postmanToMail(i))

	}
	go func() {
		wg.Wait()
		close(mailTransferPoint)
	}()
	return mailTransferPoint

}

func postmanToMail(postmanNumber int) string {
	ptm := map[int]string{
		1: "Semeynaya",
		2: "Приглашение от друга",
		3: "Информация от автосервиса",
	}
	mail, ok := ptm[postmanNumber]
	if !ok {
		return "Лотерея"
	}
	return mail
}
