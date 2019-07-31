package main

import (
  "fmt"
  "log"
  "net/http"
  "time"
)

func echoTime(w http.ResponseWriter, r *http.Request) {
  time := time.Now().Unix()
  fmt.Fprintf(w, "{ \"message\": \"Automate all the things!\", \"timestamp\": %v }", time)
}

func main() {
  http.HandleFunc("/", echoTime)
  log.Fatal(http.ListenAndServe(":8081", nil))
}
