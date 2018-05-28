package main

import (
    "fmt"
)

type ErrNegativeSqrt float64

// define function ErrNegativeSqrt.Error()
// which implements error interface.
func (e ErrNegativeSqrt) Error() string {
    // float64(a) -> Type conversion?
    return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func Sqrt(f float64) (float64, error) {
    if f < 0 {
        return 0, ErrNegativeSqrt(f)
    } else {
        return f, nil
    }
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
