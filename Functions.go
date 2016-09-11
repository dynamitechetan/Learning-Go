package main
import "fmt"

func plus(a int, b int) int {
    return a + b
}
// when having same datatype type for many params, it can be written only once 
func plusPlus(a, b, c int) int {
    return a + b + c
}

func main() {

//examples for calling a function
    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)
}
