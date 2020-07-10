// myMath.go - myMath module
package myMath
import ("errors"; "math")

func Add(args ...int) int {
    result := 0
    for _, arg := range args {
        result += arg
    }
    return result
}

func Sub(a, b int) int { return a - b }

func Mul(a, b int) int { return a * b }

func Div(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("Divide by zero illegal")
    }
    return a / b, nil
}

func Pow(a, b float64) float64 {
    return math.Pow(a, b)
}

