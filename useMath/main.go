// main.go - Main program
package main
import (
    "fmt"
    v1 "github.com/pablo0615/myMath"
    v2 "github.com/pablo0615/myMath/v2"
    )

func main() {
    add := v2.Add(1, 2, 3, 4, 5, 6)
    fmt.Println(add)
    if div, err := v1.Div(40, 5); err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(div)
    }
    pow := v1.Pow(3, 4)
    fmt.Println(pow)
}

/*****************************************

    $ go run main.go
    21 
    8
    81

*/
