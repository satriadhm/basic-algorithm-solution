package main

import "fmt"

func kabisat(tahun int) bool {
	var kondisi bool
	if tahun%400 == 0 || (tahun%4 == 0 && tahun%100 != 0) {
		kondisi = true
	}
	return kondisi
}

func angkaBulan(bulan string) int {
	var antrean int

	if bulan == "januari" {
		antrean = 1
	} else if bulan == "februari" {
		antrean = 2
	} else if bulan == "maret" {
		antrean = 3
	} else if bulan == "april" {
		antrean = 4
	} else if bulan == "mei" {
		antrean = 5
	} else if bulan == "juni" {
		antrean = 6
	} else if bulan == "juli" {
		antrean = 7
	} else if bulan == "agustus" {
		antrean = 8
	} else if bulan == "september" {
		antrean = 9
	} else if bulan == "oktober" {
		antrean = 10
	} else if bulan == "november" {
		antrean = 11
	} else if bulan == "desember" {
		antrean = 12
	} else {
		antrean = 0
	}
	return antrean
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
	bln := bulanAngka(bulan)
	kbst := kabisat(tahun)

	if (bln == "desember") || (bln == "oktober") || (bln == "agustus") || (bln == "juli") || (bln == "mei") || (bln == "maret") || (bln == "januari") {
		hari = 31
	} else if (bln == "november") || (bln == "september") || (bln == "juni") || (bln == "april") {
		hari = 30
	} else if bln == "februari" && !kbst {
		hari = 28
	} else if bln == "februari" && kbst {
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
	if hari1 == "jumat" || hari1 == "kamis" {
		*durasi = 4
	} else {
		*durasi = 2
	}
}

func pengambilan(tanggal1, bulan1, tahun1 int, hari1 string, tanggal2, bulan2, tahun2 *int, hari2 *string) {
	var day2 string
	var duration int

	mencariDurasi(hari1, &day2, &duration)
	*hari2 = day2
	*bulan2 = bulan1
	*tanggal2 = tanggal1 + duration
	*tahun2 = tahun1

	if tanggal1+duration > jumlahHari(bulan1, tahun1) {
		*tanggal2 = (tanggal1 + duration) - jumlahHari(bulan1, tahun1)
		*bulan2 = bulan1 + 1

		if bulan1+1 == 13 {
			*bulan2 = 1
			*tahun2 = tahun1 + 1
		}
	}
}

func main() {
	var hari, bulan, Hari string
	var tanggal, Tanggal, bln, tahun, Tahun int

	fmt.Scan(&hari, &tanggal, &bulan, &tahun)
	bulan1 := angkaBulan(bulan)
	pengambilan(tanggal, bulan1, tahun, hari, &Tanggal, &bln, &Tahun, &Hari)
	bulan2 := bulanAngka(bln)
	fmt.Print(Hari, " ", Tanggal, " ", bulan2, " ", Tahun)
}
