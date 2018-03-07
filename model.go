package main

type Product struct {
  ID          string     `json:"id,omitempty"`
  Name        string     `json:"name,omitempty"`
  Category    string     `json:"category,omitempty"`
  Location    *Location  `json:"location,omitempty"`
}
type Location struct {
  Store  string     `json:"store,omitempty"`
  Price  float32    `json:"price,omitempty"`
}

var inventory []Product
