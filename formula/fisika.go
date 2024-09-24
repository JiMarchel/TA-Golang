package formula

import "fmt"

// HitungDaya menghitung daya
func HitungDaya() {
	var usaha, waktu float64
	fmt.Print("Masukkan usaha (Joule): ")
	fmt.Scan(&usaha)

	fmt.Print("Masukkan waktu (detik): ")
	fmt.Scan(&waktu)

	if waktu <= 0 {
		fmt.Println("Waktu harus lebih besar dari 0.")
		return
	}

	daya := usaha / waktu
	fmt.Printf("Daya: %.2f Watt\n", daya)
}

// HitungTekanan menghitung tekanan
func HitungTekanan() {
	var gaya, luas float64
	fmt.Print("Masukkan gaya (Newton): ")
	fmt.Scan(&gaya)

	fmt.Print("Masukkan luas (m²): ")
	fmt.Scan(&luas)

	if luas <= 0 {
		fmt.Println("Luas harus lebih besar dari 0.")
		return
	}

	tekanan := gaya / luas
	fmt.Printf("Tekanan: %.2f Pascal\n", tekanan)
}

// HitungUsaha menghitung usaha
func HitungUsaha() {
	var gaya, jarak float64
	fmt.Print("Masukkan gaya (Newton): ")
	fmt.Scan(&gaya)

	fmt.Print("Masukkan jarak (meter): ")
	fmt.Scan(&jarak)

	usaha := gaya * jarak
	fmt.Printf("Usaha: %.2f Joule\n", usaha)
}

// HitungGaya menghitung gaya
func HitungGaya() {
	var massa, percepatan float64
	fmt.Print("Masukkan massa (kg): ")
	fmt.Scan(&massa)

	fmt.Print("Masukkan percepatan (m/s²): ")
	fmt.Scan(&percepatan)

	gaya := massa * percepatan
	fmt.Printf("Gaya: %.2f Newton\n", gaya)
}
