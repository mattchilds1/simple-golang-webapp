package main

import (
    "net/http"
    "log"
    "os"
    "fmt"

    "github.com/gorilla/mux"
)


func init() {
  var err error
  // var cred *mgo.Credential - for use with credentials
  // DB, err = newMongoDB("localhost", cred)
  DB, err = newMongoDB(os.Getenv("DB_URL"))  // "localhost"
  fmt.Println("DB URL:", os.Getenv("DB_URL"))
  if err != nil {
  log.Fatal(err)
  }
}

// Routes
func main() {
  r := mux.NewRouter()

  r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
  r.HandleFunc("/", indexPage)
  r.HandleFunc("/reports/", reportListPage)

  log.Fatal(http.ListenAndServe(":8080", r))
}
