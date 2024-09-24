package formula

import "fmt"

func HitungGLB() {
	var kecepatan, waktu, jarak float64

	fmt.Println("Menghitung Gerak Lurus Beraturan (GLB)")
	fmt.Print("Masukkan kecepatan (m/s): ")
	fmt.Scan(&kecepatan)
	fmt.Print("Masukkan waktu (detik): ")
	fmt.Scan(&waktu)

	jarak = kecepatan * waktu
	fmt.Printf("Jarak tempuh: %.2f meter\n", jarak)
}
