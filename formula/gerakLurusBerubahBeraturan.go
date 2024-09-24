package formula

import "fmt"

func HitungGLBB() {
	var kecepatan_awal, percepatan, waktu, jarak float64

	fmt.Println("Menghitung Gerak Lurus Berubah Beraturan (GLBB)")
	fmt.Print("Masukkan kecepatan awal (m/s): ")
	fmt.Scan(&kecepatan_awal)
	fmt.Print("Masukkan percepatan (m/s^2): ")
	fmt.Scan(&percepatan)
	fmt.Print("Masukkan waktu (detik): ")
	fmt.Scan(&waktu)

	jarak = (kecepatan_awal * waktu) + (0.5 * percepatan * waktu * waktu)
	fmt.Printf("Jarak tempuh: %.2f meter\n", jarak)
}
