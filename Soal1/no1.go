package main

import "fmt"

func main() {
	var angka []int
	for {
		var masukan int
		fmt.Scan(&masukan)
		if masukan < 0 {
			break
		}
		angka = append(angka, masukan)
	}

	insertionSort(angka)
	fmt.Println(angka)

	if cekJarak(angka) {
		fmt.Println("Data berjarak", angka[1]-angka[0])
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}

func insertionSort(angka []int) {
	for i := 1; i < len(angka); i++ {
		kunci := angka[i]
		j := i - 1
		for j >= 0 && angka[j] > kunci {
			angka[j+1] = angka[j]
			j--
		}
		angka[j+1] = kunci
	}
}

func cekJarak(angka []int) bool {
	if len(angka) < 2 {
		return true
	}
	jarak := angka[1] - angka[0]
	for i := 2; i < len(angka); i++ {
		if angka[i]-angka[i-1] != jarak {
			return false
		}
	}
	return true
}
