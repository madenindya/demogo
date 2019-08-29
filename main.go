package main

import (
	"encoding/json"
	"fmt"
)

// basics
func main() {

	// pointers
	var b *int
	fmt.Println(b) // nil
	var val int
	fmt.Println(val) // 0
	a := &val
	*a = 2
	fmt.Println(a, *a, val) // 0x00.. 2 2

	// struct
	var err error
	buku := Book{Name: "kafka on the shore", Qty: 10}
	bukuJSON, err := json.Marshal(buku)
	if err != nil {
		fmt.Println("ERROR MARSHAL!")
		return
	}
	fmt.Println(string(bukuJSON))
	// Quiz! How to remove author from result?

	// methods
	fmt.Println()
	buku.Sold(2)
	bukuJSON, _ = json.Marshal(buku)
	fmt.Println(string(bukuJSON))

	Sold2(buku, 2)
	bukuJSON, _ = json.Marshal(buku)
	fmt.Println(string(bukuJSON))

	buku.Sold3(2)
	bukuJSON, _ = json.Marshal(buku)
	fmt.Println(string(bukuJSON))

	Sold4(&buku, 2)
	bukuJSON, _ = json.Marshal(buku)
	fmt.Println(string(bukuJSON))

	// interface
	fmt.Println()
	buku = Book{Name: "The Hobbit", Qty: 10}
	shirt := Shirt{Name: "UT Uniqlo x Kaws", Qty: 0}

	CheckQty(&buku)
	CheckQty(&shirt)

	product := NewShirt(Shirt{Name: "Kaos oblong", Qty: 1})
	fmt.Println(product.Stocks())
	product = NewBook(Book{Name: "Speech and Language Processing", Qty: 2})
	fmt.Println(product.Stocks())

	// asert
	fmt.Println()
	var interfacebuku interface{} = Book{Name: "bukuku"}
	t, _ := interfacebuku.(Book)
	fmt.Println(t.Name)

}
