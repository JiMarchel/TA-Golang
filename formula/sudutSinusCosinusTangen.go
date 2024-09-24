package formula

import (
	"fmt"
	"math"
)

func HitungTrigonometri() {
	var degree float64
	fmt.Print("Masukkan nilai sudut (dalam derajat): ")
	fmt.Scan(&degree)

	radian := degree * math.Pi / 180

	sinus := math.Sin(radian)
	cosinus := math.Cos(radian)
	tangen := math.Tan(radian)

	fmt.Printf("Sinus dari %.2f derajat: %.2f\n", degree, sinus)
	fmt.Printf("Cosinus dari %.2f derajat: %.2f\n", degree, cosinus)
	fmt.Printf("Tangen dari %.2f derajat: %.2f\n", degree, tangen)
}
