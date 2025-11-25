package main

import "fmt"

func main() {

	var a int
	fmt.Scan(&a)

	switch {
	case a%10 == 0:
		fmt.Println("Kategori: Bilangan Kelipatan 10")
		fmt.Println("Hasil pembagian antara ", a, " / 10 = ", a/10)
	case a%5 == 0 && a > 5:
		fmt.Println("Kategori: Bilangan Kelipatan 5")
		fmt.Println("Hasil Kuadrat dari ", a, " ^2 = ", a*a)
	case a%2 == 1:
		fmt.Println("Kategori: Bilangan ganjil")
		fmt.Println("Hasil penjumlahan dengan bilangan berikutnya ", a, " + ", a+1, " = ", a*2+1)
	case a%2 == 0:
		fmt.Println("Kategori: Bilangan Genap")
		fmt.Println("Hasil perkalian dengan bilangan berikutnya ", a, " * ", a+1, " = ", a*a+a)
	}
}
