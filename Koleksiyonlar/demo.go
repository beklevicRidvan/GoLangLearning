package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("==== KAFE SİPARİŞ YÖNETİMİ ===")

	var tables [3]string = [3]string{"Masa 1", "Masa 2", "Masa 3"}

	fmt.Println("Masalar:", tables)

	tables[1] = "Bahçe Masası"

	fmt.Println("Güncellenmiş Masalar:", tables)

	menu := map[string]float64{
		"Kahve":   54.99,
		"Çay":     25.0,
		"Pasta":   45.0,
		"Su":      10.0,
		"Sandviç": 70.0,
	}
	fmt.Println("Menü Listesi:")
	for item, price := range menu {
		fmt.Printf("- %s: %.2f$\n", item, price)
	}

	menu["Kek"] = 30.0
	fmt.Println("Yeni ürün Eklendi: Kek - 30 TL")

	orders := []string{}

	fmt.Println("Sipariş alınıyor....")
	orders = append(orders, "Kahve", "Pasta", "Su")

	fmt.Println("Aktif Siparişler:", orders)

	orders = append(orders, "Çay")

	fmt.Println("Yeni bir sipariş eklendi:", orders)

	orders = removeAt(orders, 2) // Su silinsin
	fmt.Println("Su iptal edildi", orders)

	var total float64
	for _, item := range orders {
		price, ok := menu[item]
		if ok {
			total += price
		} else {
			fmt.Println("Menüde olmayan ürün:", item)
		}
	}
	fmt.Printf("Toplam Tutar: %2.f $ \n", total)

	customers := map[string]map[string]string{
		"123": {
			"name":     "Rıdvan",
			"lastName": "Bekleviç",
			"masa":     "Masa 1",
		},
		"456": {
			"name":     "Mehmet",
			"lastName": "Öz",
			"masa":     "Bahçe Masası",
		},
		"789": {
			"name":     "Ali",
			"lastName": "Özen",
			"masa":     "Masa 3",
		},
	}

	fmt.Println("Müşteri Listesi:")

	for id, info := range customers {
		if slices.Contains(tables[:], info["masa"]) {
			fmt.Printf("- %s: %s %s (%s)\n", id, info["name"], info["lastName"], info["masa"])
		} else {
			fmt.Println("Masa Bulunamadı.")

		}
	}

	fmt.Println("Sipraiş yönetimi tamamnlandı.")

}

// T bir type parametresidir, herhangi bir tip olabilir.
// any ise T'nin herhangi bir tip olabileceğini belirtir.
// Bu sayede fonksiyon her türlü slice ile çalışabilir.
// YAni string, int, float64 gibi tiplerle kullanılabilir.
func removeAt[T any](s []T, idx int) []T {
	// s []T: slinecek elemanın bulunduğu slice
	// idx int: silinecek elemanın indexi
	// []T: Dönüş değeri - Yeni slice (orijinal değişmez)

	// Güvenlik Kontrolü
	if idx < 0 || idx >= len(s) {
		return s // Geçersiz index ise orijinal slice'i döndür
	}

	return append(s[:idx], s[idx+1:]...)
}
