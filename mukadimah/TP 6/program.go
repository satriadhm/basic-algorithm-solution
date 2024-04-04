package main

import "fmt"

func main() {
	var hari, bulan, Hari string
	var tanggal, tahun, Tanggal, Bulan, Tahun int
	fmt.Scan(&hari, &tanggal, &bulan, &tahun)
	bulan1 := angkaBulan(bulan)
	pengambilan(tanggal, bulan1, tahun, hari, &Tanggal, &Bulan, &Tahun, &Hari)
	bulan2 := bulanAngka(Bulan)
	fmt.Print(Hari, " ", Tanggal, " ", bulan2, " ", Tahun)

}
func kabisat(tahun int) bool {
	var status bool
	if tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0) {
		status = true
	}
	return status
}
func angkaBulan(bulan string) int {
	var urutan int
	if bulan == "januari" {
		urutan = 1
	} else if bulan == "februari" {
		urutan = 2
	} else if bulan == "maret" {
		urutan = 3
	} else if bulan == "april" {
		urutan = 4
	} else if bulan == "mei" {
		urutan = 5
	} else if bulan == "juni" {
		urutan = 6
	} else if bulan == "juli" {
		urutan = 7
	} else if bulan == "agustus" {
		urutan = 8
	} else if bulan == "september" {
		urutan = 9
	} else if bulan == "oktober" {
		urutan = 10
	} else if bulan == "november" {
		urutan = 11
	} else if bulan == "desember" {
		urutan = 12
	} else {
		urutan = 0
	}
	return urutan
}
func bulanAngka(angka int) string {
	var bulan string
	if angka == 1 {
		bulan = "januari"
	} else if angka == 2 {
		bulan = "februari"
	} else if angka == 3 {
		bulan = "maret"
	} else if angka == 4 {
		bulan = "april"
	} else if angka == 5 {
		bulan = "mei"
	} else if angka == 6 {
		bulan = "juni"
	} else if angka == 7 {
		bulan = "juli"
	} else if angka == 8 {
		bulan = "agustus"
	} else if angka == 9 {
		bulan = "september"
	} else if angka == 10 {
		bulan = "oktober"
	} else if angka == 11 {
		bulan = "november"
	} else if angka == 12 {
		bulan = "desember"
	}
	return bulan
}
func jumlahHari(bulan, tahun int) int {
	var hari int
	Kabisat := kabisat(tahun)
	Bulan := bulanAngka(bulan)
	if (Bulan == "januari") || (Bulan == "maret") || (Bulan == "mei") || (Bulan == "juli") || (Bulan == "agustus") || (Bulan == "oktober") || (Bulan == "desember") {
		hari = 31
	} else if (Bulan == "april") || (Bulan == "juni") || (Bulan == "september") || (Bulan == "november") {
		hari = 30
	} else if Bulan == "februari" && !Kabisat {
		hari = 28
	} else if Bulan == "februari" && Kabisat {
		hari = 29
	}
	return hari
}
func mencariDurasi(hari1 string, hari2 *string, durasi *int) {
	if hari1 == "senin" {
		*hari2 = "rabu"
	} else if hari1 == "selasa" {
		*hari2 = "kamis"
	} else if hari1 == "rabu" {
		*hari2 = "jumat"
	} else if hari1 == "kamis" {
		*hari2 = "senin"
	} else if hari1 == "jumat" {
		*hari2 = "selasa"
	}
	if hari1 == "kamis" || hari1 == "jumat" {
		*durasi = 4
	} else {
		*durasi = 2
	}
}
func pengambilan(tanggal1, bulan1, tahun1 int, hari1 string, tanggal2, bulan2, tahun2 *int, hari2 *string) {
	var durr int
	var har2 string
	mencariDurasi(hari1, &har2, &durr)
	*hari2 = har2

	*tanggal2 = tanggal1 + durr
	*bulan2 = bulan1
	*tahun2 = tahun1
	if tanggal1+durr > jumlahHari(bulan1, tahun1) {
		*tanggal2 = (tanggal1 + durr) - jumlahHari(bulan1, tahun1)
		*bulan2 = bulan1 + 1
		if bulan1+1 == 13 {
			*bulan2 = 1
			*tahun2 = tahun1 + 1
		}
	}
}
