package formula

import "fmt"

// KonversiSuhu mengkonversi suhu dari satuan yang diberikan ke satuan lainnya
func KonversiSuhu() {
	var suhu float64
	var satuan string

	fmt.Print("Masukkan suhu: ")
	fmt.Scan(&suhu)

	fmt.Print("Masukkan satuan (C untuk Celsius, F untuk Fahrenheit, K untuk Kelvin): ")
	fmt.Scan(&satuan)

	switch satuan {
	case "C", "c":
		konversiCelsius(suhu)
	case "F", "f":
		konversiFahrenheit(suhu)
	case "K", "k":
		konversiKelvin(suhu)
	default:
		fmt.Println("Satuan tidak valid. Silakan masukkan C, F, atau K.")
	}
}

func konversiCelsius(suhu float64) {
	// Menghitung konversi dari Celsius
	fahrenheit := (suhu * 9 / 5) + 32
	kelvin := suhu + 273.15
	fmt.Printf("%.2f °C = %.2f °F\n", suhu, fahrenheit)
	fmt.Printf("%.2f °C = %.2f K\n", suhu, kelvin)
}

func konversiFahrenheit(suhu float64) {
	// Menghitung konversi dari Fahrenheit
	celsius := (suhu - 32) * 5 / 9
	kelvin := celsius + 273.15
	fmt.Printf("%.2f °F = %.2f °C\n", suhu, celsius)
	fmt.Printf("%.2f °F = %.2f K\n", suhu, kelvin)
}

func konversiKelvin(suhu float64) {
	// Menghitung konversi dari Kelvin
	celsius := suhu - 273.15
	fahrenheit := (celsius * 9 / 5) + 32
	fmt.Printf("%.2f K = %.2f °C\n", suhu, celsius)
	fmt.Printf("%.2f K = %.2f °F\n", suhu, fahrenheit)
}
