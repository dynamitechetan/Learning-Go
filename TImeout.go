// Timeouts are important for programs that connect to external resources or that otherwise need to bound execution time
package main

import "time"
import "fmt"

    

func main() {

// call that returns its result on a channel c1 after 2s. 

    c1 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c1 <- "result 1"
    }()


// If we allow a longer timeout of 3s, then the receive from c2 will succeed and we’ll print the result.
    

    c2 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "result 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(time.Second * 3):
        fmt.Println("timeout 2")
    }
}
