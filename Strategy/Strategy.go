package main

import "fmt"

type MemberStrategy interface {
	Quote(bookPrice float32) float32
}

type PrimaryMember struct {
	name string
	age  int
}

type MiddleMember struct {
	name string
	age  int
}

type HighMember struct {
	name string
	age  int
}

type Price struct {
	strategy MemberStrategy
}

func NewPrice(strategy MemberStrategy) Price {
	price := Price{}
	price.strategy = strategy
	return price
}

func (p Price) getPrice(bookPrice float32) float32 {
	return p.strategy.Quote(bookPrice)
}

func (p PrimaryMember) Quote(bookPrice float32) float32 {
	fmt.Println("primary member has no discount")
	return bookPrice
}

func (m MiddleMember) Quote(bookPrice float32) float32 {
	fmt.Println("middle member has 10% discount")
	return 0.9 * bookPrice
}

func (h HighMember) Quote(bookPrice float32) float32 {
	fmt.Println("high member has 20% discount")
	return 0.8 * bookPrice
}

func main() {
	highMember := HighMember{
		name: "Bob",
		age:  18,
	}
	price := Price{}
	price.strategy = highMember
	bookPrice := 55.89
	fmt.Printf("the books' price is: %.2f \n", price.getPrice(float32(bookPrice)))

	primaryMember := PrimaryMember{"Alice", 21}
	price.strategy = primaryMember
	fmt.Printf("the books' price is: %.2f \n", price.getPrice(float32(bookPrice)))

	middleMember := MiddleMember{"Jack", 12}
	price.strategy = middleMember
	fmt.Printf("the books' price is: %.2f \n", price.getPrice(float32(bookPrice)))
}
