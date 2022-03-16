package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Soal 1
type segitigaSamaSisi struct{
	alas, tinggi int
  }

  func (s segitigaSamaSisi) keliling() int{
	  return s.alas * 3
  }

  func (s segitigaSamaSisi) luas() int{
	  return s.alas * s.tinggi / 2
  }
  
  type persegiPanjang struct{
	panjang, lebar int
  }

  func (p persegiPanjang) keliling() int{
	  return 2 * (p.panjang + p.lebar)
  }

  func (p persegiPanjang) luas() int{
	  return p.panjang * p.lebar
  }
  
  type tabung struct{
	jariJari, tinggi float64
  }

  func (t tabung) volume() float64{
	  return math.Pi * math.Pow(t.jariJari,2) * t.tinggi
  }

  func (t tabung) luasPermukaan() float64{
	  return (2 * math.Pi * t.jariJari * t.tinggi) + (2 * math.Pi * math.Pow(t.jariJari,2))
  }
  
  type balok struct{
	panjang, lebar, tinggi int
  }

  func (b balok) volume() float64{
	  return float64(b.panjang * b.lebar * b.tinggi)
  }

  func (b balok) luasPermukaan() float64{
	return float64(2 * (b.lebar * b.panjang + b.panjang * b.tinggi + b.lebar * b.tinggi))
}

  
  type hitungBangunDatar interface{
	luas() int
	keliling() int
  }
  
  type hitungBangunRuang interface{
	volume() float64
	luasPermukaan() float64
  }

	// Soal 2
	type phone struct{
		name, brand string
		year int
		colors []string
	 }

	 type displayPhone interface{
		 display() string
	 }

	 func (p phone) display()string{
		fmt.Println("name : ", p.name)
		fmt.Println("brand : ",p.brand)
		fmt.Println("year : ",p.year)
		fmt.Println("colors : ",strings.Join(p.colors,", "))
		return ""
	 }

	//  Soal 3 return interface kosong
	func luasPersegi(sisi int, input bool) interface{}{
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
	
	
func main(){
	// Soal 1
	{
		var bangunDatar1 hitungBangunDatar
		var bangunDatar2 hitungBangunDatar
		var bangunRuang1 hitungBangunRuang
		var bangunRuang2 hitungBangunRuang

		// persegi panjang
		bangunDatar1 = persegiPanjang{10,20}
		fmt.Println("===== persegi panjang")
		fmt.Println("luas      :", bangunDatar1.luas())
		fmt.Println("keliling  :", bangunDatar1.keliling())	

		// segitiga sama sisi
		bangunDatar2 = segitigaSamaSisi{20,6}
		fmt.Println("===== segitiga sama sisi")
		fmt.Println("luas      :", bangunDatar2.luas())
		fmt.Println("keliling  :", bangunDatar2.keliling())	

		// balok
		bangunRuang1 = balok{20,100,20}
		fmt.Println("===== balok")
		fmt.Println("luas permukaan :", bangunRuang1.luasPermukaan())
		fmt.Println("volume         :", bangunRuang1.volume())	

		// tabung
		bangunRuang2 = tabung{7.0,25.0}
		fmt.Println("===== tabung")
		fmt.Println("luas permukaan :", bangunRuang2.luasPermukaan())
		fmt.Println("volume         :", bangunRuang2.volume())	

	}

	// Soal 2
	{
		var HP displayPhone
		color := []string{"Mystic Bronze","Mystic White","Mystic Black"}
		HP = phone{"Samsung Galaxy Note 20","Samsung Galaxy Note 20",2020,color}
		HP.display()
	}
	// Soal 3
	{
		fmt.Println(luasPersegi(4, true))

		fmt.Println(luasPersegi(8, false))

		fmt.Println(luasPersegi(0, true))

		fmt.Println(luasPersegi(0, false))
	}
	// Soal 4
	var prefix interface{}= "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6,8}

	var kumpulanAngkaKedua interface{} = []int{12,14}

	// convert to int
	var Angka1 = kumpulanAngkaPertama.([]int)
	hasil := Angka1[0]+Angka1[1]

	var Angka2 = kumpulanAngkaKedua.([]int)
	hasil2 := Angka2[0]+Angka2[1]
	hasil12 := hasil +  hasil2

	hasilString := []string{strconv.Itoa(Angka1[0]),strconv.Itoa(Angka1[1])}
	hasilString2 := []string{strconv.Itoa(Angka2[0]),strconv.Itoa(Angka2[1])}
	fmt.Println(prefix, strings.Join(hasilString, " + "), "+ "+ strings.Join(hasilString2, " + "), "=",hasil12)

}