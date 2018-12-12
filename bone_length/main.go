//Aufgabe 9 Scheme 1+2

package main

import "fmt"

type Calculator interface {
	b_length() float64
}

type Human struct {
	age            int
	gender         string
	bone_length_cm float64
}

func (h Human) b_length() float64 {
	var length float64
	if h.gender == "m" {
		length = 69.089 + 2.238*h.bone_length_cm
	} else {
		length = 61.412 + 2.317*h.bone_length_cm
	}

	if h.age >= 30 {
		length -= float64(h.age-30) * 0.06
	}
	return length
}

func main() {
	h := Human{age: 36, gender: "m", bone_length_cm: 20.0}
	h2 := Human{age: 15, gender: "m", bone_length_cm: 20.0}
	h3 := Human{age: 50, gender: "f", bone_length_cm: 15.5}

	fmt.Printf("%.2f\n", h.b_length())
	fmt.Printf("%.2f\n", h2.b_length())
	fmt.Printf("%.2f\n", h3.b_length())
}
