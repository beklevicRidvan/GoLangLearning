package main

import (
	"fmt"
	"os"
)

func main() {
	/*fmt.Println("Program başladı")

	// defer edilen kod main fonksiyonu bitmeden hemen önce çalışır
	defer fmt.Println("Program bitti")

	fmt.Println("Program çalışıyor")*/
	defer fmt.Println("Program bitti")

	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Dosya açılamadı", err)
		return
	}

	defer file.Close()

	// defer'da Last in first out muhabbet ivar son giren ilk çıkıyor  ilk close yapılır sonra program bitti işlemi
	fmt.Println("Dosya başarıyla açıldı")
}
