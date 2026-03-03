package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1) Menü (ürün -> fiyat)
	menu := map[string]float64{
		"kahve": 199.99,
		"pasta": 249.99,
		"çay":   79.99,
		"kek":   129.99,
		"salep": 159.90,
	}

	// 2) Sepet (ürün -> adet)
	cart := map[string]int{}

	// 3) Kupon kodu bir kere uygulanabilsin diye state
	couponApplied := false
	var couponCode string

	fmt.Println("KASA / SİPARİŞ SİSTEMİNE HOŞGELDİN")
	fmt.Println("--------------------------------")

	for {
		printMainMenu()

		var choice string
		fmt.Print("Seçim (1-7 veya q): ")
		fmt.Scanln(&choice)
		choice = strings.ToLower(strings.TrimSpace(choice))

		if choice == "q" || choice == "7" {
			fmt.Println("Çıkılıyor. İyi günler!")
			return
		}

		switch choice {
		case "1":
			// Ürünleri listele
			showProducts(menu)

		case "2":
			// Sepete ekle
			product, qty, ok := readProductAndQty("Eklenecek ürün adı: ", "Adet: ")
			if !ok {
				continue
			}
			ok = addToCart(cart, menu, product, qty)
			if ok {
				fmt.Println("✅ Sepete eklendi.")
			} else {
				fmt.Println("❌ Sepete eklenemedi.")
			}

		case "3":
			// Sepetten çıkar
			product, qty, ok := readProductAndQty("Çıkarılacak ürün adı: ", "Adet: ")
			if !ok {
				continue
			}
			ok = removeFromCart(cart, product, qty)
			if ok {
				fmt.Println("✅ Sepetten çıkarıldı.")
			} else {
				fmt.Println("❌ Sepetten çıkarılamadı.")
			}

		case "4":
			// Sepeti göster
			printCart(cart, menu)

		case "5":
			// Kupon uygula
			if couponApplied {
				fmt.Println("⚠️ Kupon zaten uygulandı:", couponCode)
				continue
			}

			fmt.Print("Kupon kodu (SAVE10): ")
			var code string
			fmt.Scanln(&code)
			code = strings.ToUpper(strings.TrimSpace(code))

			total := calcTotal(cart, menu)
			newTotal, ok := applyCoupon(total, code)
			if !ok {
				fmt.Println("❌ Geçersiz kupon.")
				continue
			}

			// Kuponu “state” olarak saklıyoruz
			couponApplied = true
			couponCode = code
			fmt.Printf("✅ Kupon uygulandı. Yeni toplam: %.2f TL\n", newTotal)

		case "6":
			// Ödeme
			if len(cart) == 0 {
				fmt.Println("⚠️ Sepet boş. Önce ürün ekleyin.")
				continue
			}

			total := calcTotal(cart, menu)
			if couponApplied {
				discounted, ok := applyCoupon(total, couponCode)
				if ok {
					total = discounted
				}
			}

			fmt.Printf("Ödenecek toplam: %.2f TL\n", total)
			fmt.Print("Müşteri ne kadar verdi? (örn 500): ")

			var paid float64
			fmt.Scanln(&paid)

			change, ok := checkout(total, paid)
			if !ok {
				fmt.Println("❌ Yetersiz ödeme. İşlem iptal.")
				continue
			}

			fmt.Printf("✅ Ödeme alındı. Para üstü: %.2f TL\n", change)
			fmt.Println("🧾 Fiş:")
			printReceipt(cart, menu, couponApplied, couponCode)

			// Ödeme sonrası sepet sıfırlama
			cart = map[string]int{}
			couponApplied = false
			couponCode = ""

		default:
			fmt.Println("Geçersiz seçim. 1-7 veya q girin.")
		}

		fmt.Println() // boş satır
	}
}

func printMainMenu() {
	fmt.Println("MENÜ")
	fmt.Println("1) Ürünleri listele")
	fmt.Println("2) Sepete ürün ekle")
	fmt.Println("3) Sepetten ürün çıkar")
	fmt.Println("4) Sepeti göster")
	fmt.Println("5) Kupon uygula (SAVE10)")
	fmt.Println("6) Ödeme al")
	fmt.Println("7) Çıkış")
}

// ====== CORE FUNCTIONS ======

func showProducts(menu map[string]float64) {
	fmt.Println("ÜRÜNLER")
	for name, price := range menu {
		fmt.Printf("- %s: %.2f TL\n", name, price)
	}
}

func addToCart(cart map[string]int, menu map[string]float64, product string, qty int) bool {
	if qty <= 0 {
		fmt.Println("Adet 1 veya daha büyük olmalı.")
		return false
	}
	if _, exists := menu[product]; !exists {
		fmt.Println("Menüde böyle bir ürün yok.")
		return false
	}
	cart[product] += qty
	return true
}

func removeFromCart(cart map[string]int, product string, qty int) bool {
	if qty <= 0 {
		fmt.Println("Adet 1 veya daha büyük olmalı.")
		return false
	}

	current, exists := cart[product]
	if !exists {
		fmt.Println("Sepette böyle bir ürün yok.")
		return false
	}

	if qty > current {
		fmt.Println("Sepette o kadar ürün yok.")
		return false
	}

	if qty == current {
		delete(cart, product)
	} else {
		cart[product] = current - qty
	}
	return true
}

func calcTotal(cart map[string]int, menu map[string]float64) float64 {
	var total float64
	for product, qty := range cart {
		price := menu[product] // burada product’ın menu’de olduğu varsayımı var (addToCart garanti ediyor)
		total += price * float64(qty)
	}
	return total
}

func printCart(cart map[string]int, menu map[string]float64) {
	if len(cart) == 0 {
		fmt.Println("Sepet boş.")
		return
	}

	fmt.Println("SEPET")
	var subtotal float64
	for product, qty := range cart {
		price := menu[product]
		lineTotal := price * float64(qty)
		subtotal += lineTotal
		fmt.Printf("- %s x%d = %.2f TL\n", product, qty, lineTotal)
	}
	fmt.Printf("Ara toplam: %.2f TL\n", subtotal)
}

func applyCoupon(total float64, code string) (float64, bool) {
	// Basit kupon: SAVE10 => %10 indirim
	if code != "SAVE10" {
		return total, false
	}
	discounted := total * 0.90
	return discounted, true
}

func checkout(total float64, paid float64) (change float64, ok bool) {
	if paid < total {
		return 0, false
	}
	return paid - total, true
}

// ====== INPUT HELPERS ======

func readProductAndQty(productPrompt, qtyPrompt string) (product string, qty int, ok bool) {
	fmt.Print(productPrompt)
	fmt.Scanln(&product)
	product = strings.ToLower(strings.TrimSpace(product))

	fmt.Print(qtyPrompt)
	fmt.Scanln(&qty)

	if product == "" {
		fmt.Println("Ürün adı boş olamaz.")
		return "", 0, false
	}
	if qty <= 0 {
		fmt.Println("Adet 1 veya daha büyük olmalı.")
		return "", 0, false
	}
	return product, qty, true
}

// ====== RECEIPT ======

func printReceipt(cart map[string]int, menu map[string]float64, couponApplied bool, couponCode string) {
	var total float64
	for product, qty := range cart {
		price := menu[product]
		lineTotal := price * float64(qty)
		total += lineTotal
		fmt.Printf("%s x%d  %.2f TL\n", product, qty, lineTotal)
	}

	fmt.Printf("Ara toplam: %.2f TL\n", total)

	if couponApplied {
		discounted, ok := applyCoupon(total, couponCode)
		if ok {
			fmt.Printf("Kupon (%s) uygulandı.\n", couponCode)
			fmt.Printf("İndirimli toplam: %.2f TL\n", discounted)
			return
		}
	}

	fmt.Printf("Toplam: %.2f TL\n", total)
}
