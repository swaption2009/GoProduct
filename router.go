package main

import (
  "github.com/gorilla/mux"
)

func registerRouter() *mux.Router {
  router := mux.NewRouter()
  router.HandleFunc("/inventories", GetInventoryEndpoint).Methods("GET")
  router.HandleFunc("/inventories", CreateProductEndpoint).Methods("POST")
  router.HandleFunc("/inventories/{id}", GetProductEndpoint).Methods("GET")
  // TODO add UpdateProductHandler
  // router.HandleFunc("/inventories/{id}", UpdateProductHandler).Methods("PUT")
  router.HandleFunc("/inventories/{id}", DeleteProductEndpoint).Methods("DELETE")

  return router
}
