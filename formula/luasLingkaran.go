package formula

import (
	"fmt"
	"math"
)

func LuasLingkaran() {
	var radius float64
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&radius)

	luas := math.Pi * radius * radius
	fmt.Printf("Luas lingkaran dengan jari-jari %.2f adalah %.2f\n", radius, luas)
}
