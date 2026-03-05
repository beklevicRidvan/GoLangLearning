package main

import "fmt"

func main() {

	// Slice Header için bellek ayırır
	// ptr,len , cap
	// ptr = nil pointer,
	// Elemanlar için Hiç bellek ayrılmaz

	ps := new([]int)

	fmt.Println("Ps Adresi", ps)
	fmt.Println("Ps slice'i nil mi ?", *ps == nil)

	// append çalışır mı // çalışır

	*ps = append(*ps, 1, 5, 5)
	fmt.Println("Slice'in durumu: ", *ps)
	fmt.Println("len:", len(*ps), "cap:", cap(*ps))

}
