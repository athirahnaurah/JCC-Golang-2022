package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Soal 1
	{
		var word1 = "Jabar"
		var word2 = "Coding"
		var word3 = "Camp"
		var word4 = "Golang"
		var word5 = "2022"
	
		fmt.Println(word1 +" "+ word2 +" "+ word3 + " "+ word4 +" "+ word5)
	}

	// Soal 2
	{
		halo := "Halo Dunia"
		halo2 := strings.Replace(halo,"Dunia","Golang",1)

		fmt.Println(halo2)
	}

	// Soal3
	{
		var kataPertama = "saya"
		var kataKedua = "senang"
		var kataKetiga = "belajar"
		var kataKeempat = "golang"

		var kataKedua2 = strings.Title(kataKedua)
		var kataKetiga2 = strings.Replace(kataKetiga,"r","R",1)
		var kataKeempat2 = strings.ToUpper(kataKeempat)

		fmt.Println(kataPertama +" "+ kataKedua2 + " " + kataKetiga2 + " " + kataKeempat2)
	}

	// Soal 4
	{
		var angkaPertama= "8"
		var angkaKedua= "5"
		var angkaKetiga= "6"
		var angkaKeempat = "7"

		angkaPertama2, _:= strconv.Atoi(angkaPertama)
		angkaKedua2, _ := strconv.Atoi(angkaKedua)
		angkaKetiga2, _ := strconv.Atoi(angkaKetiga)
		angkaKeempat2, _ := strconv.Atoi(angkaKeempat)

		hasil := angkaPertama2 + angkaKedua2 + angkaKetiga2 + angkaKeempat2

		fmt.Println(hasil)
	}

	// Soal 5
	{
		kalimat := "halo halo bandung"
		angka := 2022

		kalimat2 := strings.Replace(kalimat, "halo","Hi",2)
		fmt.Printf("%q",kalimat2)
		fmt.Printf(`-`)
		fmt.Print(angka)
	}	
}