package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	min := 1
	max := 20
	randomNum := rand.IntN(max-min) + min
	fmt.Println(randomNum)
}
