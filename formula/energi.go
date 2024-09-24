package formula

import "fmt"

// Energi potensial: Ep = m * g * h
func HitungEnergiPotensial() {
	var massa, gravitasi, tinggi float64
	fmt.Print("Masukkan massa (kg): ")
	fmt.Scan(&massa)
	fmt.Print("Masukkan gravitasi (m/s²): ")
	fmt.Scan(&gravitasi)
	fmt.Print("Masukkan tinggi (m): ")
	fmt.Scan(&tinggi)

	energiPotensial := massa * gravitasi * tinggi
	fmt.Printf("Energi Potensial: %.2f Joule\n", energiPotensial)
}

// Energi kinetik: Ek = 1/2 * m * v²
func HitungEnergiKinetik() {
	var massa, kecepatan float64
	fmt.Print("Masukkan massa (kg): ")
	fmt.Scan(&massa)
	fmt.Print("Masukkan kecepatan (m/s): ")
	fmt.Scan(&kecepatan)

	energiKinetik := 0.5 * massa * kecepatan * kecepatan
	fmt.Printf("Energi Kinetik: %.2f Joule\n", energiKinetik)
}
