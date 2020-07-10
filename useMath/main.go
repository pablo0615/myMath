// main.go - Main program
package main
import (
    "fmt"
    gm "github.com/pablo0615/myMath"
    gmv2 "github.com/pablo0615/myMath/v2"
    )

func main() {
    add := gmv2.Add(1, 2, 3, 4, 5, 6)
    fmt.Println(add)
    if div, err := gm.Div(40, 5); err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(div)
    }
    pow := gm.Pow(3, 4)
    fmt.Println(pow)
}

/*****************************************

    $ go run main.go
    21 
    8
    81

*/
