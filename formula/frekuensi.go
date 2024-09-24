package formula

import "fmt"

// Frekuensi: f = 1 / T
func HitungFrekuensi() {
	var periode float64
	fmt.Print("Masukkan periode (detik): ")
	fmt.Scan(&periode)

	if periode <= 0 {
		fmt.Println("Periode harus lebih besar dari 0.")
		return
	}

	frekuensi := 1 / periode
	fmt.Printf("Frekuensi: %.2f Hertz\n", frekuensi)
}
