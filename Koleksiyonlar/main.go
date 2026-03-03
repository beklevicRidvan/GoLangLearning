package main

import "fmt"

func main() {

	// Array => sabit boyutlu listelerdir -> Belirli sayıda elemen tutar
	// tek bir veri tipi tutarlar
	// Value typetır. Kopyalanır referans tutulamaz.
	// Boyutu sonradan değiştirilemez
	// Yeni eleman eklenemez
	// Mecut eleman silinemez
	// Eleman güncellenebilir.

	var numbers [3]int = [3]int{10, 20, 30}
	fmt.Println("Array Elemanları: ", numbers[0], numbers[1], numbers[2])

	numbers[1] = 99

	fmt.Println("Güncelleme sonrası array: ", numbers)

	fmt.Println("------------------------------------------")
	fmt.Println("------------------------------------------")

	// array de aralık seçimi - slicing
	subArray := numbers[0:2]
	// 0. indexten başlar 2. indexe kadar elemanları alır.
	// 2. index dahil değil.
	fmt.Println("Alt array:", subArray)

	subArray2 := numbers[1:]
	fmt.Println("Alt array 2:", subArray2)

	subArray3 := numbers[:2]
	fmt.Println("Alt arrray 3: ", subArray3)

	fmt.Println("------------------------------------------")
	fmt.Println("------------------------------------------")

	// İç içe array örenği - matrix

	var matrix [2][3]int = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println("0.boyut 1.index: ", matrix[0][1])
	fmt.Println("1.boyut 2.index: ", matrix[1][2])

	// Slice
	// Dinamik uzunluktadır. Boyutu sonradan değiştirilebilir.
	// Tek bir veri tipi tutarlar
	// Hafızada bir array referans tutarlar.
	// Reference typetır. Kopyalanamaz , referans tutulabilir.
	// Gerçek hayatta da genellikle array den çok slicelar kullanılır.

	var numberSlice []int = []int{10, 20, 30}
	fmt.Println("Slice Elemanları: ", numberSlice[0], numberSlice[1], numberSlice[2])

	// eleman ekleme
	numberSlice = append(numberSlice, 40)
	fmt.Println("Eklendikten sonra slice: ", numberSlice)

	// birden fazla eleman ekleme
	moreNumbers := []int{50, 60}
	numberSlice = append(numberSlice, moreNumbers...)
	fmt.Println("Daha fazla eleman eklendikten sonra: ", numberSlice)

	// güncelleme
	numberSlice[0] = 45
	fmt.Println("Güncelleme sonrası:", numberSlice)

	// Slice aralık seçme slicing
	subSlice := numberSlice[1:4]
	// 1. indexten başlar 4. index e kadar
	// 4. index dahil değil!
	fmt.Println("Alt slice:", subSlice)

	// 2. indexten başla sonuna kadar git
	subSlice2 := numberSlice[2:]
	fmt.Println("Alt slice 2:", subSlice2)

	subSlice3 := numberSlice[:4]
	// 0. indexten başla 4'e kadar git. 4. index dahil değil
	fmt.Println("Alt slice3: ", subSlice3)

	// Slice silme
	// 2. indexteki elemanı sildirme
	indexToRemove := 2
	fmt.Println("2. indexe kadar olan elemanlar:", numberSlice[:indexToRemove])
	fmt.Println("2. indexten sonraki elemanlar:", numberSlice[indexToRemove+1:])
	numberSlice = append(numberSlice[:indexToRemove], numberSlice[indexToRemove+1:]...)
	fmt.Println("Eleman silindikten sonra:", numberSlice)

	matrixSlice := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("0.boyut 1.index: ", matrixSlice[0][1])
	fmt.Println("1.boyut 2.index: ", matrixSlice[1][2])

	fmt.Println("------------------------------------------")
	fmt.Println("------------------------------------------")

	// Map
	// Anahtar - Değer key - value
	// hızlı veri erişimi sağlarlar
	// Anahtar benzersizdir unique
	// Değerler herhangi bir veri tipi olabilir

	prices := map[string]float64{
		"Kahve": 54.90,
		"Çay":   25.90,
		"Pasta": 149.99,
	}
	fmt.Println("Kahve Fiyatı: ", prices["Kahve"])

	prices["Kek"] = 30.0
	fmt.Println("Güncel Liste: ", prices)

	prices["Pasta"] = 300.56
	fmt.Println("Güncel Liste: ", prices)

	delete(prices, "Çay")
	fmt.Println("Güncel Liste: ", prices)

	// İç içe map

	persons := map[string]map[string]string{
		"123": {
			"name":     "Rıdvan",
			"lastName": "Bekleviç",
		},
		"456": {
			"name":     "Mehmet",
			"lastName": "Öz",
		},
		"789": {
			"name":     "Ali",
			"lastName": "Özen",
		},
	}

	fmt.Println("Person 123 adı: ", persons["123"]["name"])
	fmt.Println("Person 456 soyadı: ", persons["456"]["lastName"])

	// Map kontrol etme

	price, exists := prices["Pasta"]
	if exists {
		fmt.Println("Pasta fiyatı: ", price)
	} else {
		fmt.Println("Pasta bulunamadı")
	}
	fmt.Println("------------------------------------------")
	fmt.Println("------------------------------------------")
	fmt.Println("------------------------------------------")
	fmt.Println("------------------------------------------")

	// Döngüler

	fmt.Println("Numbers Arrayinin eleman sayısı", len(numbers))
	for i := 0; i < len(numbers); i++ {
		fmt.Println("Index", i, "Değer", numbers[i])
	}

	// Array üzerinde range ile döngü kurma

	for i, v := range numbers {
		fmt.Println("Index", i, "Değer", v)
	}

	for i, v := range numberSlice {
		fmt.Println("Index", i, "Değer", v)
	}

	for k, v := range prices {
		fmt.Println(k, "Fiyat", v)

	}

	fmt.Println("map sadece keyler")
	for k := range prices {
		fmt.Println("Anahtar", k)
	}

	for _, v := range prices {
		fmt.Println("Değer: ", v)
	}

}
