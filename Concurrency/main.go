package main

import (
	"fmt"
	"time"
)

// Eşzamanlılık
//

func main() {
	slowJob("1.işlem")
	slowJob("2.işlem")

	// Goroutine nedi?
	// Normal bir fonksiyon tek başına sırayla çalışır.(senkron)
	// Goroutine ise fonksiyonu arka planda asenkron çalıştırır (Future.wait[])

	go slowJob("3. işlem")
	go slowJob("4. işlem")

	// Bunu koymamızın nedeni arkaplanda görmüyoruz bekleten bişey olması lazım
	time.Sleep(3 * time.Second)

}

func slowJob(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(name, "-> adım", i)
		time.Sleep(500 * time.Millisecond)
	}
}
