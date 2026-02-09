package main

import "fmt"

const DAYS_FOR_PAYMENT = 5

func main() {
	// GoodsCount:=5
	// goods_count:=5
	// goods4you:=5
	goodsCount := 5
	fmt.Println(goodsCount, DAYS_FOR_PAYMENT)
	// лучше создать новую переменнную
	goodsCount = -8
	fmt.Println(goodsCount)
}
