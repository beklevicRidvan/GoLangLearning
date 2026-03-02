package main

import "fmt"

func main() {
	fmt.Println("Gün Sonu Kafe Raporu")
	fmt.Println("---------------------")

	coffeePrice := 140.90
	cakePrice := 225.90
	coffeeSold := 37
	caleSold := 20

	totalCoffeRevenue := float64(coffeeSold) * coffeePrice

	totalCakeRevenue := float64(caleSold) * cakePrice

	totalRevenue := totalCoffeRevenue + totalCakeRevenue

	fmt.Printf("Kahve Geliri: %.2f TL\n", totalCoffeRevenue)
	fmt.Printf("Pasta Geliri: %.2f TL\n", totalCakeRevenue)
	fmt.Printf("Toplam Gelir: %.2f TL\n", totalRevenue)

	fmt.Println("------------------------")

	rent := 1250.0
	rent += 150.0 // Elektrik faturası
	rent += 200.0 // Su faturası
	rent += 300.0 // Malzeme giderleri

	fmt.Printf("Giderler: %.2f TL\n", rent)

	netProfit := totalRevenue - rent
	fmt.Printf("Net Kar: %.2f TL\n", netProfit)

	isWeekendCampaign := true
	hasMembershipCampaign := false

	fmt.Println("Hafta sonu ve üyelere özel kampanya var mı?", isWeekendCampaign && hasMembershipCampaign)
	fmt.Println("Hafta sonu veya üyelere özel kampanya var mı?", isWeekendCampaign || hasMembershipCampaign)

	if isWeekendCampaign {
		fmt.Println("Hafta sonu kampanyası aktif!")
	}

	// kasadaki bozuk para sayma

	coint := 23
	reminder := coins % 5

	fmt.Printf("Bozuk para sayısı: %d, 5'e bölümünden kalan: %d\n", coint, reminder)
}
