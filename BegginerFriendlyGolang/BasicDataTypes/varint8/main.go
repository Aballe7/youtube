package main

import (
        "fmt"
)

func main() {
        //int8: integers (-128 to 127)
        var a int8 = 100 + 100
        fmt.Println(a) // constant 200 overflows int8
}
