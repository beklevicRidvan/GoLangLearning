package main

import (
	"fmt"
	"sync"
	"time"
)

func sendMessageFunc(ch chan string, msg string, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(1 * time.Second)
	ch <- msg

}

func main() {
	ch := make(chan string)
	var wg sync.WaitGroup

	// 3 tane goroutine çalıştır
	wg.Add(3)

	go sendMessageFunc(ch, "Mesaj 1", &wg)
	go sendMessageFunc(ch, "Mesaj 2", &wg)
	go sendMessageFunc(ch, "Mesaj 3", &wg)

	// Gözetmen goroutine oluşturuyoruz -> işçiler bitince kanalı kapatacak
	go func() {
		wg.Wait() // 3 goroutine tammamlanana kadar bekle
		close(ch) // kanalı kapat

	}()

	// Tüm mesajları dinle
	for mesaj := range ch {
		fmt.Println("Aldım: ", mesaj)

	}

}
