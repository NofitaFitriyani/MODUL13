package main

import "fmt"

type Buku struct {
	ID        int
	Judul     string
	Penulis   string
	Penerbit  string
	Eksemplar int
	Tahun     int
	Rating    int
}

type DaftarBuku []Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Scanln(n)
	*pustaka = make(DaftarBuku, *n)
	for i := 0; i < *n; i++ {
		fmt.Scanln(&(*pustaka)[i].ID)
		fmt.Scanln(&(*pustaka)[i].Judul)
		fmt.Scanln(&(*pustaka)[i].Penulis)
		fmt.Scanln(&(*pustaka)[i].Penerbit)
		fmt.Scanln(&(*pustaka)[i].Eksemplar)
		fmt.Scanln(&(*pustaka)[i].Tahun)
		fmt.Scanln(&(*pustaka)[i].Rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku) {
	maxIndex := 0
	for i := 1; i < len(pustaka); i++ {
		if pustaka[i].Rating > pustaka[maxIndex].Rating {
			maxIndex = i
		}
	}
	buku := pustaka[maxIndex]
	fmt.Println(buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun)
}

func UrutBuku(pustaka *DaftarBuku) {
	for i := 0; i < len(*pustaka)-1; i++ {
		for j := 0; j < len(*pustaka)-i-1; j++ {
			if (*pustaka)[j].Rating < (*pustaka)[j+1].Rating {
				(*pustaka)[j], (*pustaka)[j+1] = (*pustaka)[j+1], (*pustaka)[j]
			}
		}
	}
}

func Cetak5Terbaru(pustaka DaftarBuku) {
	batas := 5
	if len(pustaka) < 5 {
		batas = len(pustaka)
	}
	for i := 0; i < batas; i++ {
		buku := pustaka[i]
		fmt.Println(buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun)
	}
}

func CariBuku(pustaka DaftarBuku, r int) {
	low, high := 0, len(pustaka)-1
	found := false
	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].Rating == r {
			found = true
			buku := pustaka[mid]
			fmt.Println(buku.Judul, buku.Penulis, buku.Penerbit, buku.Eksemplar, buku.Tahun, buku.Rating)
			break
		} else if pustaka[mid].Rating < r {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if !found {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var pustaka DaftarBuku
	var n, rating int

	// Pendaftaran buku
	fmt.Println("Masukkan jumlah buku dan data buku:")
	DaftarkanBuku(&pustaka, &n)

	// Cetak buku dengan rating tertinggi
	fmt.Println("Buku dengan rating tertinggi:")
	CetakTerfavorit(pustaka)

	// Urutkan buku berdasarkan rating
	UrutBuku(&pustaka)

	// Cetak 5 buku terbaru berdasarkan rating
	fmt.Println("5 Buku dengan rating tertinggi:")
	Cetak5Terbaru(pustaka)

	// Cari buku berdasarkan rating tertentu
	fmt.Println("Masukkan rating yang ingin dicari:")
	fmt.Scanln(&rating)
	fmt.Println("Hasil pencarian buku dengan rating:")
	CariBuku(pustaka, rating)
}
