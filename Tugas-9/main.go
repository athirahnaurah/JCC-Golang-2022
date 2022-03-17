package main

import (
	"Tugas-9/library"
	"fmt"
	"strconv"
	"strings"
)
func main(){
	// Soal 1
	{
		var bangunDatar1 library.HitungBangunDatar
		var bangunDatar2 library.HitungBangunDatar
		var bangunRuang1 library.HitungBangunRuang
		var bangunRuang2 library.HitungBangunRuang

		// persegi panjang
		bangunDatar1 = library.PersegiPanjang{Panjang:10,Lebar:20}
		fmt.Println("===== persegi panjang")
		fmt.Println("luas      :", bangunDatar1.Luas())
		fmt.Println("keliling  :", bangunDatar1.Keliling())	

		// segitiga sama sisi
		bangunDatar2 = library.SegitigaSamaSisi{Alas:20,Tinggi:6}
		fmt.Println("===== segitiga sama sisi")
		fmt.Println("luas      :", bangunDatar2.Luas())
		fmt.Println("keliling  :", bangunDatar2.Keliling())	

		// balok
		bangunRuang1 = library.Balok{Panjang:20,Lebar:100,Tinggi:20}
		fmt.Println("===== balok")
		fmt.Println("luas permukaan :", bangunRuang1.LuasPermukaan())
		fmt.Println("volume         :", bangunRuang1.Volume())	

		// tabung
		bangunRuang2 = library.Tabung{JariJari:7.0,Tinggi:25.0}
		fmt.Println("===== tabung")
		fmt.Println("luas permukaan :", bangunRuang2.LuasPermukaan())
		fmt.Println("volume         :", bangunRuang2.Volume())	

	}

	// Soal 2
	{
		var HP library.DisplayPhone
		color := []string{"Mystic Bronze","Mystic White","Mystic Black"}
		HP = library.Phone{Name:"Samsung Galaxy Note 20",Brand:"Samsung Galaxy Note 20",Year:2020,Colors:color}
		HP.Display()
	}
	// Soal 3
	{
		fmt.Println(library.LuasPersegi(4, true))

		fmt.Println(library.LuasPersegi(8, false))

		fmt.Println(library.LuasPersegi(0, true))

		fmt.Println(library.LuasPersegi(0, false))
	}
	
	// Soal 4
	{
		var angka1 = library.KumpulanAngkaPertama.([]int)
		hasil := angka1[0]+angka1[1]
	
		var angka2 = library.KumpulanAngkaKedua.([]int)
		hasil2 := angka2[0]+angka2[1]
	
		hasil12 := hasil +  hasil2
	
		hasilString := []string{strconv.Itoa(angka1[0]),strconv.Itoa(angka1[1])}
		hasilString2 := []string{strconv.Itoa(angka2[0]),strconv.Itoa(angka2[1])}
	
		fmt.Println(library.Prefix, strings.Join(hasilString, " + "), "+ "+ strings.Join(hasilString2, " + "), "=",hasil12)
	}
	
}