package main
	
import "fmt"
	
type dude struct {
    name string
    age  int
}

func main() {
// one way to create a struct
    fmt.Println(dude{"Chetan", 18})	
// another way to create a struct
    fmt.Println(dude{name: "Dynamite", age: 4})

    fmt.Println(dude{name: "Pappu"})  // age is 0 (default)

// Access struct

    s := dude{name: "Alan", age: 25}
    fmt.Println(s.name)

// struct with pointers
    sp := &s
    fmt.Println(sp.age)

// Structs are mutable.
    sp.age = 51
    fmt.Println(sp.age)
}
