package main

import "fmt"
func main() {
// channel of strings buffering up to 2 values.
messages := make(chan string, 2)
messages <- "buffered"
messages <- "channel"
  // retreival of values
  fmt.Println(<-messages)
  fmt.Println(<-messages)
}
