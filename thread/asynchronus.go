package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UTC().UnixNano())

  query := "Our Query"
  respond := make(chan string)

  go googleIt(respond, query)

  fmt.Println("before respond")  	
  //queryResp := <-respond
  fmt.Println("after respond")  	

  fmt.Printf("Sent query:\t\t %s\n", query)
  //fmt.Printf("Got Response:\t\t %s\n", queryResp)
}

func googleIt(respond chan<- string, query string) {
  time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
  fmt.Println("googleIt") 
  respond <- "A Google Response"
}
