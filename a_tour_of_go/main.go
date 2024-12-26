package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	const Epsilon = .000001
	z := 1.0
	for i := 0; i < 10; i++ {
		zNew := z - (z*z - x) / (2*z)
		fmt.Println("z = ", z)
		if math.Abs(zNew - z) < Epsilon {
			break
		}
		z = zNew
	}
	return z
}

func main() {
	fmt.Println(Sqrt((2)))
	p1 := Person { "Heiko", 62 }
	p2 := Person { "Doris", 66 }

	fmt.Println(p1)
	fmt.Println(p2)

	PrintAge(p1)

}

func PrintAge(aging Aging) {
	fmt.Println("AGE = ", aging.Age())
}

type Aging interface {
	Age() int
}

type Person struct {
	name string
	age int
}

func (p Person) String() string {
	return fmt.Sprintf("Person: Name=%s, age=%d\n", p.name, p.age)
}

func (p Person) Age() int {
	return p.age
}