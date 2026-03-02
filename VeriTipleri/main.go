package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Go statik tipli bir dildir.
	// Yani her değişken tipi derleme (compile time) zamanında belirlenir.

	// Hatalar kod çalıştırılmadan önce tespit edillir.
	// Kod okunabilirliğini artar.
	// Performans yükü azalır.

	// int8 tipi -128 ile 127 arasında değer alabilir 		  				1 byte yer kaplar.
	// int16 tipi -32,768 ile 32,767 arasında değer alabilir  				2 byte yer kaplar.
	// int32 tipi -2,147,483,648 ile 2,147,483,647 arasında değer alabilir  4 byte yer kaplar.
	// int64 tipi -9,223,372,036,854,775,808 ile 9,223,372,036,854,775,807 arasında değer alabilir
	// 																		8 byte yer kaplar.
	// int tipi ise sistem mimarisine bağlı olarak 32 bit veya 64 bit olabilir.
	// 															 Bu da göre 4 veya 8 byte yer kaplar.

	// Optimizasyon gerekmedikçe int tipi kullanılması önerilir.

	var age int = 30
	fmt.Println(age)

	var age2 int16 = 40
	fmt.Println(age2)

	// uint veri tipi sadece pozitif tam sayıları tutar. 0 ile başlayarak pozitif tam sayıları içerir.
	var positiveAge uint16 = 25
	fmt.Println(positiveAge)

	// float32 ve float64 veri tipleri ondalık sayıları tutar. float32 4 byte, float64 ise 8 byte yer kaplar.
	var price float32 = 19.99
	fmt.Println(price)

	var pi float64 = 3.14159
	fmt.Printf("Pi sayısı: %.2f\n", pi)

	// string veri tipi metinleri tutar. Stringler çift tırnak içinde tanımlanır."
	var name string = "Rıdvan"
	fmt.Println(name)

	// bool veri tipi true veya false değerlerini tutar.
	var isStudent bool = true
	fmt.Println(isStudent)

	var x int = 13
	var y float64 = float64(x) // int tipini float64'e dönüştürme
	fmt.Printf("x: %d, y: %.2f\n", x, y)

	var z float64 = 9.99
	var w int = int(z) // float64 tipini int'e dönüştürme
	fmt.Printf("z: %.2f, w: %d\n", z, w)

	// int to string
	var num int = 100
	var sayi string = fmt.Sprintf("%d", num)
	fmt.Println(sayi)

	// string to int
	var sayi2 string = "200"
	var num2 int
	fmt.Sscanf(sayi2, "%d", &num2)
	fmt.Println(num2)

	// strconv paketi ile dönüşümler yapılabilir

	// string to int
	// Atoi'nin anlamı "ASCII to Integer" demektir.(ASCII (string ifadeler))
	numStr := "300"
	num3, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Println("Sayı:", num3)
	}

	// int to string
	num4 := 400
	numStr2 := strconv.Itoa(num4) // Itoa'nın anlamı "Integer to ASCII" demektir.
	fmt.Println("Sayı String:", numStr2)

	//float64'ten string'e dönüşüm
	floatVal2 := 2.7986
	flatStr2 := strconv.FormatFloat(floatVal2, 'f', 2, 64) // 'f' ondalık format, 2 ondalık basamak, 64 bit float

	fmt.Println("Float String:", flatStr2)
}
