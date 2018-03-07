package main

import "fmt"

func SeedData() {
  inventory = append(inventory, Product{ID: "1", Name: "Banana", Category: "Fruits", Location: &Location{Store: "Target", Price: 1.00}})
  inventory = append(inventory, Product{ID: "2", Name: "Coconut Water", Category: "Drinks", Location: &Location{Store: "Meijer", Price: 5.95}})

  fmt.Println("2 records created!")
}