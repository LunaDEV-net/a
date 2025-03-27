package main

import (
	"fmt"
	"math"
	"time"
)

func calc_t(nachkommaStellen int) int {
	vV := 0.0
	vH := -1.0
	mV := 1.0
	mH := math.Pow(100, float64(nachkommaStellen))
	c := 0
	
	for {
		oldVV := vV
		oldVH := vH
		
		vH = (mH*oldVH + mV*(2*oldVV-oldVH)) / (mV + mH)
		vV = (mV*oldVV + mH*(2*oldVH-oldVV)) / (mV + mH)
		c++
		
		if vV < 0 {
			vV = -vV
			c++
		}
		
		if vH > 0 && vH > math.Abs(vV) {
			break
		}
	}
	return c
}

func main() {
	// Check if the user provided an argument
	for i := 1; i < 10+1; i++ {
		gestartet := time.Now()
		fmt.Printf("Starting @%s\n", gestartet)
		
		result := calc_t(i)
		
		fmt.Printf("Finished @%s\n", time.Now())
		fmt.Printf("Berechnungsdauer: %s\n", time.Since(gestartet))
		fmt.Printf("Berechnete Nachkommastellen: %d\n", i)
		fmt.Printf("Anzahl der Kollisionen: %d\n", result)
		fmt.Printf("Unser berechnetes π:   %.15f\n", float64(result)/math.Pow(10, float64(i)))
		fmt.Printf("Das echte π:           %.15f\n", math.Pi)
		fmt.Println()
		fmt.Println("----")
		fmt.Println()
	}
	
}
