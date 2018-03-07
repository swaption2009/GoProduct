package main

import (
  "log"
  "net/http"
)

func main() {
  SeedData()

  router := registerRouter()
  log.Fatal(http.ListenAndServe(":8000", router))
}
