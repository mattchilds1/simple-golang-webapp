package main

import (
    "net/http"
    "log"

    "github.com/gorilla/mux"
)


func init() {
  var err error
  // var cred *mgo.Credential - for use with credentials
  // DB, err = newMongoDB("localhost", cred)
  DB, err = newMongoDB("localhost")

  if err != nil {
  log.Fatal(err)
  }
}

// Routes
func main() {
  r := mux.NewRouter()

  r.HandleFunc("/", indexPage)
  r.HandleFunc("/reports/", reportListPage)

  log.Fatal(http.ListenAndServe(":8080", r))
}
