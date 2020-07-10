// main.go - Main program
package main
import (
    "fmt"
    gm "github.com/pablo0615/myMath"
    )

func main() {
    add := gm.Add(1, 2)
    fmt.Println(add)
    if div, err := gm.Div(40, 5); err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(div)
    }
}

/*****************************************

    $ go run main.go
    3
    8

*/
