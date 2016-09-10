package main
import "fmt"
func main() {	

    nums := []int{2, 3, 4}
    sum := 0
// this for loop will add the elements of nums to variable sum
    for _, num := range nums {
        sum += num
    }
    fmt.Println(sum)
//this for loop will print index of element "3" in nums array
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
// this will print unicode values
    for i, c := range "go" {
        fmt.Println(i, c)
    }
}
