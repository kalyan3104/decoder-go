package main

import (
	"fmt"
)


func calculateCGPA(sgpa []float64) float64 {
	if len(sgpa) == 0 {
		return 0.0
	}

	var total float64
	for _, s := range sgpa {
		total += s
	}

	cgpa := total / float64(len(sgpa))
	return cgpa
}

func main() {
	
	sgpa := []float64{6.8, 8.3, 6.8, 8.17}

	cgpa := calculateCGPA(sgpa)

	// fmt.Printf("The CGPA is: %.2f\n", cgpa)
	fmt.Println(cgpa)
}
