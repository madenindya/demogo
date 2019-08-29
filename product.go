package main

import (
	"fmt"
	"strconv"
)

// Product ....
type Product interface {
	Stocks() int
	IsAvailable() bool
}

// CheckQty ...
func CheckQty(p Product) {
	fmt.Println("Amount: " + strconv.Itoa(p.Stocks()) + " left!")
}
