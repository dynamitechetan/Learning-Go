package main

import "fmt"

func main() {	
    messages := make(chan string)	
    go func() { messages <- "hi" }()	
    msg := <-messages
    fmt.Println(msg)
}
