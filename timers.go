package main

import "time"
import "fmt"	

func main() {
// This timer will wait 2 seconds.
    timer1 := time.NewTimer(time.Second * 2)
    <-timer1.C
    fmt.Println("Timer 1 expired")


}
