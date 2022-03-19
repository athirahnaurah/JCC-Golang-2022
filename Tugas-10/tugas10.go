package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"strconv"
	"time"
)

// Soal 1
func development(sentence string, year int){
	kalimat := sentence
	tahun := year
	fmt.Println(kalimat, tahun)
}

// Soal 2
func kelilingSegitigaSamaSisi(value int, input bool)(string, error){
	sisi := value
	var sentence string
	var err error
	defer endApp(err)
	if sisi == 0 && input == false{
		err = errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
	}else if sisi == 0 && input == true{
		err = errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
	}else if sisi >0 && input == false{
		sentence = strconv.Itoa(sisi)
	}else if sisi > 0 && input == true {
		sisiStr := strconv.Itoa(sisi)
		kel := 3*sisi
		kelStr := strconv.Itoa(kel)
		sentence = "keliling segitiga sama sisinya dengan sisi "+ sisiStr +" cm adalah "+ kelStr+ " cm"
	}

	if err != nil{
		panic(err)
	}
	return sentence,err
}

func endApp(err error){
	message := recover()
	fmt.Println("Error:", message)
  }

//   Soal 3
func tambahAngka(a int, b *int) (*int) {
	*b += a
	return (b)
}

func cetakAngka(a *int){
	fmt.Println(*a)
}

// Soal 4
func addPhone (phone *[]string, merek... string){
	*phone = append(*phone,merek...)
}

// Soal 5
func kelilingLingkaran(r float64)float64{
	kel := math.Pi * 2 * r 
	return math.Round(kel)
}

func luasLingkaran(r float64) float64 {
	luas := math.Pi * math.Pow(r, 2)
	return math.Round(luas)
	
  }

// Soal 6
func kelilingPersegiPanjang(p int64,l int64)int64{
	return (2*p)+(2*l) 
}

func luasPersegiPanjang(p int64,l int64)int64{
	return p*l
}

func main(){
	// Soal 1
	{
		defer development("Golang Backend Development",2021)
		fmt.Println("Mulai run program")
	}
	// Soal 2
	{
		fmt.Println(kelilingSegitigaSamaSisi(4, true))

		fmt.Println(kelilingSegitigaSamaSisi(8, false))

		fmt.Println(kelilingSegitigaSamaSisi(0, true))

		fmt.Println(kelilingSegitigaSamaSisi(0, false))

	}
	// Soal 3
	{
		angka := 1
		defer cetakAngka(&angka)

		tambahAngka(7, &angka)

		tambahAngka(6, &angka)

		tambahAngka(-1, &angka)

		tambahAngka(9, &angka)
	}
	// Soal 4
	{
		var phones = []string{}
		addPhone(&phones, "Xiaomi","Asus","iPhone","Samsung","Oppo","Realme","Vivo")
		for i, phone := range phones {
			fmt.Printf("%d. %s\n", i+1, phone)
			time.Sleep(time.Second * 1)
		}
	}
	// Soal 5
	{
		fmt.Println("Keliling Lingkaran: ",kelilingLingkaran(7),"\nLuas Lingkaran:",luasLingkaran(7))
		fmt.Println("Keliling Lingkaran: ",kelilingLingkaran(10),"\nLuas Lingkaran:",luasLingkaran(10))
		fmt.Println("Keliling Lingkaran: ",kelilingLingkaran(15),"\nLuas Lingkaran:",luasLingkaran(15))
	}
	// Soal 6
	{
		var panjang = flag.Int64("panjang", 12, "panjang persegi panjang")
		var lebar = flag.Int64("lebar", 10, "lebar persegi panjang")

		flag.Parse()
		fmt.Println("Keliling Persegi: ",kelilingPersegiPanjang(*panjang,*lebar))
		fmt.Println("Luas Persegi: ",luasPersegiPanjang(*panjang,*lebar))
	}
}
