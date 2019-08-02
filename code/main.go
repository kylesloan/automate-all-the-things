package main

// packages to use https://medium.com/golangspec/import-declarations-in-go-8de0fd3ae8ff
import (
  "fmt"
  "log"
  "net/http"
  "time"
  "os"
)

// get the hostname so we can add it to a http header to see that we are getting load balanced
func getName() string {
  name, err := os.Hostname()
  if err != nil { panic(err) }
  return name
}

// http object with now timestamp
func echoTime(w http.ResponseWriter, r *http.Request) {
  time := time.Now().Unix()
  w.Header().Set("ServedBy", getName())
  fmt.Fprintf(w, "{ \"message\": \"Automate all the things!\", \"timestamp\": %v }", time)
}

func main() {
  http.HandleFunc("/", echoTime)
  log.Fatal(http.ListenAndServe(":8081", nil))
}
