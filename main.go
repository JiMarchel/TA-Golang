package main

import (
	"bisaai_ta/formula"
	"fmt"
)

func main() {
	var selected_formula int

	fmt.Println("Halo👋, Selamat Datang")
	fmt.Println("Piih rumus apa yang mau dihitung:")

	for {
		fmt.Println("1.Menghitung luas persegi \n2.Menghitung luas segitiga \n3.Menghitung luas lingkaran \n4.Menghitung sudut sinus,cosinus,tangen \n5.Meghitung gerak lurus beraturan \n6.Menghitung gerak lurus berubah beraturan \n7.Hitung energi potensial \n8.Hitung energi kinetik \n9.Menghitung frekuensi \n10.Menghitung masa jenis \n11.Menghitung daya \n12.Menghitung tekanan \n13.Menghitung usaha \n14.Menghitung gaya \n15.Konversi satuan suhu")
		fmt.Scan(&selected_formula)

		switch selected_formula {
		case 1:
			formula.LuasPersegi()
			return
		case 2:
			formula.LuasSegitiga()
			return
		case 3:
			formula.LuasLingkaran()
			return
		case 4:
			formula.HitungTrigonometri()
			return
		case 5:
			formula.HitungGLB()
			return
		case 6:
			formula.HitungGLBB()
			return
		case 7:
			formula.HitungEnergiPotensial()
			return
		case 8:
			formula.HitungEnergiKinetik()
			return
		case 9:
			formula.HitungFrekuensi()
			return
		case 10:
			formula.HitungMasaJenis()
			return
		case 11:
			formula.HitungDaya()
			return
		case 12:
			formula.HitungTekanan()
			return
		case 13:
			formula.HitungUsaha()
			return
		case 14:
			formula.HitungGaya()
			return
		case 15:
			formula.KonversiSuhu()
			return
		default:
			fmt.Println("\nTolong pilih salah satu dari nomor 1 sampai 10🫵👊😡😡")
		}
	}
}
