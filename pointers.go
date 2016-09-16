package main

import "fmt"

func zerovalue(val int) {
    val = 0
}
func zeroptr(ptr *int) {
    *ptr = 0
}
func main() {
    i := 1
    fmt.Println("initial:", i)
    zerovalue(i)
    fmt.Println("zerovalue:", i)
//The &i syntax gives the memory address of i.
    zeroptr(&i)
    fmt.Println("zeroptr:", i)  
    fmt.Println("pointer address:", &i)
}
