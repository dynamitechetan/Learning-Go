// slice type is an abstraction built on top of Go's array type, and so to understand slices we must first understand arrays. 

package main

import "fmt"

func main() {
	
    s := make([]string, 3)
    fmt.Println("Empty : ", s)	
//basics  ... set,get,length just like arrays
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("Set:", s)
    fmt.Println("Get:", s[2])	

    fmt.Println("Length:", len(s))	

//the code below will append d,e,f to s
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)	
//the code below copies s to c
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)	
}
