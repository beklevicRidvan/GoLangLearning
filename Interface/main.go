package main

import "fmt"

type Greeter interface {
	Greet()
}

type TurkishGreeter struct {
}
type EnglishGreeter struct {
}

func (t TurkishGreeter) Greet() {
	fmt.Println("Merhaba")
}

func (e EnglishGreeter) Greet() {
	fmt.Println("Hello")
}

func SayHello(g Greeter) {
	fmt.Println("Selamlama başladı")
	g.Greet()
}

func main() {

	tr := TurkishGreeter{}
	en := EnglishGreeter{}

	SayHello(tr)
	SayHello(en)

}
