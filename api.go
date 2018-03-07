package main

import (
  "encoding/json"
  "github.com/gorilla/mux"
  "net/http"
  "github.com/satori/go.uuid"
)

func GetInventoryEndpoint(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(inventory)
}

func CreateProductEndpoint(w http.ResponseWriter, r *http.Request) {
  var product Product
  _ = json.NewDecoder(r.Body).Decode(&product)
  product.ID = uuid.Must(uuid.NewV4()).String()
  inventory = append(inventory, product)
  json.NewEncoder(w).Encode(inventory)
}

func GetProductEndpoint(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  for _, product := range inventory {
    if product.ID == params["id"] {
      json.NewEncoder(w).Encode(product)
      return
    }
  }
  json.NewEncoder(w).Encode(&Product{})
}

// TODO fix UpdateProductEndpoint
// func UpdateProductEndpoint(w http.ResponseWriter, r *http.Request) {

// }

func DeleteProductEndpoint(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  for index, item := range inventory {
    if item.ID == params["id"] {
      inventory = append(inventory[:index], inventory[index+1:]...)
      break
    }
    json.NewEncoder(w).Encode(inventory)
  }
}
