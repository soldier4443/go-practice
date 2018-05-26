package main

// value receiver and pointer receiver
// -> immutable and mutable property

import (
	"fmt"
	"math"
)

type Integer int32
type MutableInteger int32

func (i Integer) tenth(n int32) int32 {
	return int32(i) * int32(math.Pow(10.0, float64(n)))
}

func (i *MutableInteger) tenth(n int32) int32 {
	*i = MutableInteger(int32(*i) *
		int32(math.Pow(10.0, float64(n))))

	return int32(*i)
}

func main() {
	val := Integer(30)

	fmt.Println("Immutable..")
	fmt.Println(val)
	fmt.Println(val.tenth(3))
	fmt.Println(val)
	fmt.Println("")

	mutableVal := MutableInteger(30)

	fmt.Println("Mutable..")
	fmt.Println(mutableVal)
	fmt.Println(mutableVal.tenth(3))
	fmt.Println(mutableVal)
}
