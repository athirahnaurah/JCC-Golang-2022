package library

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Soal 1
type SegitigaSamaSisi struct{
	Alas, Tinggi int
}

func (s SegitigaSamaSisi) Keliling() int{
	return s.Alas * 3
}

func (s SegitigaSamaSisi) Luas() int{
	return s.Alas * s.Tinggi / 2
}

type PersegiPanjang struct{
	Panjang, Lebar int
}

func (p PersegiPanjang) Keliling() int{
	return 2 * (p.Panjang + p.Lebar)
}

func (p PersegiPanjang) Luas() int{
	return p.Panjang * p.Lebar
}

type Tabung struct{
	JariJari, Tinggi float64
}

func (t Tabung) Volume() float64{
	return math.Pi * math.Pow(t.JariJari,2) * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64{
	return (2 * math.Pi * t.JariJari * t.Tinggi) + (2 * math.Pi * math.Pow(t.JariJari,2))
}

type Balok struct{
	Panjang, Lebar, Tinggi int
}

func (b Balok) Volume() float64{
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}

func (b Balok) LuasPermukaan() float64{
	return float64(2 * (b.Lebar * b.Panjang + b.Panjang * b.Tinggi + b.Lebar * b.Tinggi))
}


type HitungBangunDatar interface{
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface{
	Volume() float64
	LuasPermukaan() float64
}

	// Soal 2
type Phone struct{
	Name, Brand string
	Year int
	Colors []string
}

type DisplayPhone interface{
	Display() string
}

func (p Phone) Display()string{
	fmt.Println("name : ", p.Name)
	fmt.Println("brand : ",p.Brand)
	fmt.Println("year : ",p.Year)
	fmt.Println("colors : ",strings.Join(p.Colors,", "))
	return ""
}

	//Soal 3 return interface kosong
func LuasPersegi(sisi int, input bool) interface{}{
	var result interface{}
	luas := sisi * sisi
	if sisi == 0 && input == true {
		result = "Maaf anda belum menginput sisi dari persegi"
	}else if sisi == 0 && input == false{
		result = nil
	}else if input == true{
		result = "luas persegi dengan sisi " + strconv.Itoa(sisi) +" cm adalah " + strconv.Itoa(luas)+ " cm"
	}else if input == false{
		result = sisi
	}
	return result
}

	// Soal 4
var Prefix interface{}= "hasil penjumlahan dari "

var KumpulanAngkaPertama interface{} = []int{6,8}

var KumpulanAngkaKedua interface{} = []int{12,14}
	
	
