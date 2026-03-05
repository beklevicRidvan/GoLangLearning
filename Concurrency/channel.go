package main

import (
	"fmt"
	"time"
)

func sendMessage(ch chan string) {
	time.Sleep(1 - time.Second)

	ch <- "Selam! Gorutine'den mesaj geldi."

}

// Channel, goroutineler arasında veri alışverişini yapmak için kullanılır.
// Güvenli bir iletişim kurmak
func main() {

	ch := make(chan string)

	go sendMessage(ch)

	// Kanalı dinleme.

	msg := <-ch

	fmt.Println("Ana program mesaj aldı:", msg)

	// slice := make([]int, 0, 1000) -> 1000 eleman koyucam slice'a bana 1000 elemanlık bir alan ayır.
}
