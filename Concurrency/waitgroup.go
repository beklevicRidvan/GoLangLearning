package main

import (
	"fmt"
	"sync"
	"time"
)

// Waitgroup: Goroutinlerin tamamlanmasını beklemek için kullanılan sayaçtır.
// wg.Add() -> Kaç tane goroutine çalışacak
// wg.Done() -> Bu goroutine işini bitirdi.
// wg.Wait() -> Tüm goroutine'ler Done diyene kadar bekle
func main() {
	fmt.Println("---- Senkron Çalışma (go olmadan) ----")
	slowJob("1. İşlem")
	slowJob("2. İşlem")

	var wg sync.WaitGroup

	wg.Add(2) // 2 tane dediğimiz için hangi 2'si önce biterse onları umursuyor eğer geç biten varsa yetişemezse  diğerlerini umursamaz almaz.
	// eş zamanlı olarak en kısa süren 2 'sini esas alır diğerleri çalışırken eğer erken biterse işlem diğer veriler gelmeyebilir
	// Hangisi en kısaysa onu esas alır yani çalışma mantığı olarak
	fmt.Println("---- Asenkron Çalışma (go ile) ----")

	go slowJobAsync("3. İşlem", &wg, 500)  // 1500 ms sürer
	go slowJobAsync("4. İşlem", &wg, 800)  // 2400 ms sürer
	go slowJobAsync("5. İşlem", &wg, 900)  // 2700 ms sürer
	go slowJobAsync("6. İşlem", &wg, 1000) // 3000 ms sürer

	wg.Wait() // go işlevleri ne zaman bitimini bekler kafadan vermeyiz yani değeri
}

func slowJob(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(name, "-> adım", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func slowJobAsync(name string, wg *sync.WaitGroup, ms int) {
	defer wg.Done()

	for i := 1; i <= 3; i++ {

		time.Sleep(time.Duration(ms) * time.Millisecond)
		fmt.Println(name, "-> adım", i)
	}
}
