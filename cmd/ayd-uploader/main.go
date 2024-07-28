package main

import (
  "fmt"
  "net/http"
  "log"
)

func ask(w http.ResponseWriter, r *http.Request) {
  r.ParseMultipartForm(10 << 20)
  file, handler, err := r.FormFile("doc")

  if err != nil {
    log.Println("error retrieving the file")
    log.Fatal(err)
    return
  }

  defer file.Close() // close file when function ends

  mb := float64(handler.Size / 1048576)
  fmt.Printf("%s has %f MBs\n", handler.Filename, mb)

  fmt.Fprintf(w, "Hello there!")
}

func main() {
  http.HandleFunc("/ask", ask)
  log.Println("[ayd-uploader] up and running on port 8080")
  log.Fatal(http.ListenAndServe(":8080", nil))
}
