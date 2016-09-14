// Go supports anonymous functions, which can form closures.
package main
    
import "fmt"
// This function returns another function
func intSeq() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}

    
func main() {

    nextInt := intSeq()    
// calling the function
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())    

}
