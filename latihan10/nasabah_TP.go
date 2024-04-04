package main

import "fmt"

type nasabah struct {
	kode_asdos     string
	nama           string
	bank           string
	nomor_rekening int
}

type tabNasabah []nasabah

func addNasabah(T *tabNasabah, N *int) {
	var (
		kode_asdos, nama, bank string
		nomor_rekening         int
		nasabah_baru           nasabah
	)
	if len(*T) >= *N {
		fmt.Println("data penuh")
	} else {
		fmt.Scanln(&kode_asdos, &nama, &bank, &nomor_rekening)
		nasabah_baru.kode_asdos = kode_asdos
		nasabah_baru.nama = nama
		nasabah_baru.bank = bank
		nasabah_baru.nomor_rekening = nomor_rekening
		*T = append(*T, nasabah_baru)
	}
}

func cetak(T tabNasabah, N int, X string) {
	var i int = 0
	for i < N {
		nasabah := T[i]
		if nasabah.bank == X {
			fmt.Printf("Kode : %v, Nasabah: %v, Bank: %v, Rek: %v\n", nasabah.kode_asdos, nasabah.nama, nasabah.bank, nasabah.nomor_rekening)
		}
		i++
	}

}
func main() {
	var (
		daftar_nasabah tabNasabah
		jumlah_nasabah int
		nama_bank      string
	)
	jumlah_nasabah = 10
	for i := 0; i < 10; i++ {
		addNasabah(&daftar_nasabah, &jumlah_nasabah)
	}
	fmt.Scanln(&nama_bank)
	cetak(daftar_nasabah, jumlah_nasabah, nama_bank)
}
