package main

import "fmt"

func main() {
    var a int //use var to initialize a variable with zero value
    var b string
    var c float64 
    var d bool

    e := 3.14
    
    fmt.Printf("var a int \t %T [%v]\n", a, a)
    fmt.Printf("var b string \t %T [%v]\n", b, b)
    fmt.Printf("var c float \t %T [%v]\n", c, c)
    fmt.Printf("var d bool \t %T [%v]\n", d, d)
    
    fmt.Printf("%T %v\n", e, e)

    /* Conversion over casting
    to keep integrity, to keep our memory safe*/
    aaa := float32(10)
    fmt.Printf("print aaa: type: %T, value: %v\n", aaa, aaa)
}


