package main

import "fmt"
//this function below returns 2 values
func fucnt() (int, int) {
    return 3, 7
}

func main() {	
//calling the function
    a, b := fucnt()
    fmt.Println(a)
    fmt.Println(b)	
}
