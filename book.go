package main

// Author ...
type Author struct {
	AuthorName string `json:"author_name,omitempty"`
	AuthorAge  int    `json:"author_age,omitempty"`
}

// Book ...
type Book struct {
	Name   string  `json:"name,omitempty"`
	Qty    int     `json:"qty,omitempty"`
	Author *Author `json:"author,omitempty"`
}

// Sold p with the amount of n
func (b Book) Sold(n int) {
	b.Qty = b.Qty - n
}

// Sold2 ....
func Sold2(b Book, n int) {
	b.Qty = b.Qty - n
}

// Sold3 ....
func (b *Book) Sold3(n int) {
	b.Qty = b.Qty - n
}

// Sold4 ....
func Sold4(b *Book, n int) {
	b.Qty = b.Qty - n
}

// Quiz! Which have the same meaning?
// Quiz! Which will decrease the value of Book.Qty?

// Stocks check reminding stocks
func (b *Book) Stocks() int {
	return b.Qty
}

// IsAvailable check reminding stocks
func (b *Book) IsAvailable() bool {
	return b.Qty > 0
}

// NewBook ...
func NewBook(b Book) Product {
	return &b
}
