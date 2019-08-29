package main

// Shirt ...
type Shirt struct {
	Name string `json:"name,omitempty"`
	Qty  int    `json:"qty,omitempty"`
	Size string
}

// Stocks check reminding stocks
func (b *Shirt) Stocks() int {
	return b.Qty
}

// IsAvailable check reminding stocks
func (b *Shirt) IsAvailable() bool {
	return b.Qty > 0
}

// NewShirt ...
func NewShirt(b Shirt) Product {
	return &b
}
