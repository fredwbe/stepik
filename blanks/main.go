package main

import (
	"fmt"
	"math"
)

func main() {
	productPrice := 100.12
	discountPercent := 20.11
	discountAmount := (productPrice * discountPercent) / 100
	productPrice = productPrice - discountAmount
	fmt.Println(math.Floor(productPrice*100) / 100)
}
