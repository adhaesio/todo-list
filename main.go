package main

import (
	"fmt"
	"study/http"
	todo_app "study/todo_app"
)

// //type Message struct {
// 	Author string
// 	Text   string
// }

// var mtx = sync.Mutex{}

// var money = 1000
// var bank = 0

// type httpResponse struct {
// 	money          int
// 	paymentHistory []PaymentInfo
// }

// func payHandler(w http.ResponseWriter, r *http.Request) {
// 	var payment PaymentInfo
// 	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
// 		fmt.Println("err:", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	payment.Time = time.Now()

// 	payment.Println()

// 	mtx.Lock()
// 	if money-payment.USD >= 0 {
// 		money -= payment.USD
// 	}
// 	paymentHistory = append(paymentHistory, payment)

// 	HttpResponse := httpResponse{
// 		money:          money,
// 		paymentHistory: paymentHistory,
// 	}

// 	b, err := json.Marshal(HttpResponse)
// 	if err != nil {
// 		fmt.Println("err,", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	if _, err := w.Write(b); err != nil {
// 		fmt.Println("err", err)
// 		return
// 	}

// 	fmt.Println("Новое значение переменной money :", money)

// 	fmt.Println("История оплат :", paymentHistory)

// 	mtx.Unlock()
// }

// type PaymentInfo struct {
// 	//описание покупки
// 	Description string `json: Description`
// 	//сумма покупки
// 	USD int `json: Usd`
// 	//ِФИО ЧЕЛОВЕКА, СОВЕРШАЮЩЕГО ПОКУПКУ
// 	FullName string `json: Fullname`
// 	//ِМесто прописки человека, совершающего покупку
// 	Adress string `json: Adress`

// 	Time time.Time
// }

// func (p PaymentInfo) Println() {
// 	fmt.Println("Description", p.Description)
// 	fmt.Println("USD", p.USD)
// 	fmt.Println("FullName", p.FullName)
// 	fmt.Println("Adress", p.Adress)

// }

// //var paymentHistory = make([]PaymentInfo, 0)

func main() {

	todoList := todo_app.NewList()
	httpHandlers := http.NewHTTPHandlers(todoList)
	httpServer := http.NewHTTPServer(httpHandlers)
	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start http server", err)
	}

	//http.HandleFunc("/pay", payHandler)

	//if err := http.ListenAndServe(":9091", nil); err != nil {
	//
	//		fmt.Println("Error server", err)
	//	}

}
