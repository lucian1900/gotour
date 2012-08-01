package main

import (
	"fmt"
	"math"
)

var tolerance = 0.0001

func Sqrt(x float64) float64 {
	z := 1.0
	
	for i := 0; i < 10; i+=1 {
		new_z := z - (z*z - x) / (2 * z)
		//fmt.Println("Iteration ", i, z, new_z)
		
		if math.Abs(new_z - z) < tolerance {
			break
		} else {
			z = new_z
		}
	}
	
	if i := float64(int(z)); math.Abs(i - z) < tolerance {
		return i
	}
	return z
}

func main() {
	v := 9.0
	fmt.Println(Sqrt(v))
	fmt.Println(math.Sqrt(v))
}
