package formula

import "fmt"

// MasaJenis: ρ = m / V
func HitungMasaJenis() {
	var massa, volume float64
	fmt.Print("Masukkan massa (kg): ")
	fmt.Scan(&massa)

	fmt.Print("Masukkan volume (liter): ")
	fmt.Scan(&volume)

	if volume <= 0 {
		fmt.Println("Volume harus lebih besar dari 0.")
		return
	}

	masaJenis := massa / volume
	fmt.Printf("Masa Jenis: %.2f kg/L\n", masaJenis)
}
