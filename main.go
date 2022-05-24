package main

import (
    "fmt"
    "net/http"
    "math/rand"
    "time"

    "github.com/gorilla/mux"
)

func main() {
  fmt.Println("Starting Simple OTEL Helloworld Application")
  
  rand.Seed(time.Now().UnixNano())
  // Adding Gorilla Mux Router
  r := mux.NewRouter()
  r.HandleFunc("/", HomeHandler)

  http.ListenAndServe(":8080", r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  FakeLatency()
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("Hello World"))
}

func FakeLatency(){
   w := randInt(200, 1000)
   fmt.Printf("Faking A latency of: %d ms", w)
   time.Sleep(time.Duration(w) * time.Millisecond)
}

func randInt(min int, max int) int {
  return min + rand.Intn(max-min)
}
