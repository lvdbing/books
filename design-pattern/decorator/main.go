package main

import "fmt"

func main() {
	var b1 Beverage = &Espresso{}
	b1 = NewMocha(b1)
	b1 = NewMilk(b1)
	b1 = NewSoy(b1)
	b1 = NewWhip(b1)
	b1 = NewMocha(b1)
	fmt.Printf("%s ￥%d\n", b1.GetDescription(), b1.Cost())
	fmt.Println()

	var b2 Beverage = &DarkRoast{}
	b2 = NewMocha(b2)
	b2 = NewWhip(b2)
	fmt.Printf("%s ￥%d\n", b2.GetDescription(), b2.Cost())
	fmt.Println()
}
