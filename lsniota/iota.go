package main

import "fmt"

const (
	red = iota
	green
	blue
)

const (
	orange = iota
	yellow
	violet
)

func main() {
	fmt.Println(red, green, blue)
	fmt.Println(orange, yellow, violet)
}
