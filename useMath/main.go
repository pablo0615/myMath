// main.go - Main program
package main
import (
    "fmt"
    gm "github.com/pablo0615/myMath"
    )

func main() {
    add := gm.Add(1 ,2)
    fmt.Println(add)
    mult := gm.Div(40, 5)
    fmt.Println(mult)
}

/*****************************************

    $ go run main.go
    3
    8

*/
