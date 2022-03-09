package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Soal 1
	{
		var panjangPersegiPanjang string = "8"
		var lebarPersegiPanjang string = "5"

		var alasSegitiga string = "6"
		var tinggiSegitiga string = "7"

		panjangPersegiPanjang2, _ := strconv.Atoi(panjangPersegiPanjang)
		lebarPersegiPanjang2, _ := strconv.Atoi(lebarPersegiPanjang)

		alasSegitiga2, _:= strconv.Atoi(alasSegitiga)
		tinggiSegitiga2, _:= strconv.Atoi(tinggiSegitiga)


		var kelilingPersegiPanjang int = (2 * panjangPersegiPanjang2) + (2 * lebarPersegiPanjang2)
		var luasSegitiga int = alasSegitiga2 * tinggiSegitiga2 / 2

		fmt.Print("Keliling Persegi Panjang = ")
		fmt.Println(kelilingPersegiPanjang)

		fmt.Print("Luas Segitiga = ")
		fmt.Println(luasSegitiga)
	}

	// Soal 2
	{
		var nilaiJohn = 80
		var nilaiDoe = 50
		var indeksJohn string
		var indeksDoe string
		
		switch {
		case nilaiJohn >= 80:
			indeksJohn = "A"
		case nilaiJohn >=70 && nilaiJohn < 80:
			indeksJohn = "B"
		case nilaiJohn >=60 && nilaiJohn < 70:
			indeksJohn = "C"
		case nilaiJohn >=50 && nilaiJohn < 60:
			indeksJohn = "D"
		case nilaiJohn < 50:
			indeksJohn = "E"
		}

		fmt.Println("Nilai John: " + indeksJohn)

		switch {
		case nilaiDoe >= 80:
			indeksDoe = "A"
		case nilaiDoe >=70 && nilaiDoe < 80:
			indeksDoe = "B"
		case nilaiDoe >=60 && nilaiDoe < 70:
			indeksDoe = "C"
		case nilaiDoe >=50 && nilaiDoe < 60:
			indeksDoe = "D"
		case nilaiDoe < 50:
			indeksDoe = "E"
		}
		fmt.Println("Nilai Doe: " + indeksDoe)
	}

	// Soal 3
	{
		var tanggal = 1
		var bulan = 10 
		var tahun = 2002
		var bulanString string

		switch {
		case bulan == 1:
			bulanString = "Januari"
		case bulan == 2:
			bulanString = "Februari"
		case bulan == 3:
			bulanString = "Maret"
		case bulan == 4:
			bulanString = "April"
		case bulan == 5:
			bulanString = "Mei"
		case bulan == 6:
			bulanString = "Juni"
		case bulan == 7:
			bulanString = "Juli"
		case bulan == 8:
			bulanString = "Agustus"
		case bulan == 9:
			bulanString = "September"
		case bulan == 10:
			bulanString = "Oktober"
		case bulan == 11:
			bulanString = "November"
		case bulan == 12:
			bulanString = "Desember"
		default:
			fmt.Println("Bukan termasuk bulan")
		}

		tanggalString := strconv.Itoa(tanggal)
		tahunString := strconv.Itoa(tahun)

		fmt.Println(tanggalString + " " + bulanString + " " +tahunString)
	}

	// Soal 4
	{
		tahunLahir := 2002
		var generasi string

		if (tahunLahir >= 1944 && tahunLahir <= 1964){
			generasi = "Baby boomer"
		}else if(tahunLahir >= 1965 && tahunLahir <= 1979){
			generasi = "X"
		}else if(tahunLahir >= 1980 && tahunLahir <= 1994){
			generasi = "Y (Millenials)"
		}else if (tahunLahir >= 1995 && tahunLahir <= 2015){
			generasi = "Z"
		}else{
			fmt.Println("Tidak termasuk ke generasi manapun")
		}

		fmt.Println("Aku termasuk generasi " + generasi)
	}
}