package main

import "fmt"

// Soal 1
type buah struct {
	nama       string
	warna      string
	adaBijinya bool
	harga      int
}

func (b buah) printBuah() {
	fmt.Println("Nama : ",b.nama)
	fmt.Println("Warna : ",b.warna)
	if(b.adaBijinya == true){
		fmt.Println("Biji: Ada")
	}else{
		fmt.Println("Biji: Tidak")
	}
	fmt.Println("Harga : ",b.harga)
	fmt.Println()
}

// Soal 2
type segitiga struct{
	alas, tinggi int
  }
  
  type persegi struct{
	sisi int
  }
  
  type persegiPanjang struct{
	panjang, lebar int
  }

  func (segi3 segitiga) hitungLuasSegitiga(a int, t int) int{
	segi3.alas = a
	segi3.tinggi = t
	return segi3.alas * segi3.tinggi /2
}

func (segi4 persegi) hitungLuasPersegi(s int) int{
	segi4.sisi = s
	return segi4.sisi*segi4.sisi
}

func (rectangular persegiPanjang) hitungLuasPPanjang(p int, l int) int{
	rectangular.panjang = p
	rectangular.lebar = l

	return rectangular.panjang*rectangular.lebar 
}

// Soal 3
type phone struct{
	name, brand string
	year int
	colors []string
 }

func (p phone) coloring(color... string) []string {
	p.colors= append(p.colors, color...)
	return p.colors
}

// Soal 4
type movie struct {
	title string 
	genre string
	duration int 
	year int
}

func tambahDataFilm(judul string, jam int, jenis string, tahun int, m *[]movie) []movie{
	var film = movie{title: judul, genre: jenis, duration: jam/60,year: tahun }
	*m = append(*m, film)
	return *m
}

func main() {
	// Soal 1
	{
		var pinneaple buah
		pinneaple.nama = "nanas"
		pinneaple.warna = "Kuning"
		pinneaple.adaBijinya = false
		pinneaple.harga = 9000
		pinneaple.printBuah()

		var orange buah
		orange.nama = "Jeruk"
		orange.warna = "Oranye"
		orange.adaBijinya = true
		orange.harga = 8000
		orange.printBuah()

		var watermelon buah
		watermelon.nama = "Semangka"
		watermelon.warna = "Hijau & Merah"
		watermelon.adaBijinya = true
		watermelon.harga = 10000
		watermelon.printBuah()

		var banana buah
		banana.nama = "Pisang"
		banana.warna = "Kuning"
		banana.adaBijinya = false
		banana.harga = 5000
		banana.printBuah()
		fmt.Println("--------------------------------")
	}

	// Soal 2
	{
		var luas segitiga
		var luas2 persegi
		var luas3 persegiPanjang
		fmt.Println("L Segitiga: ", luas.hitungLuasSegitiga(5,10))
		fmt.Println("L Persegi: ", luas2.hitungLuasPersegi(9))
		fmt.Println("L Persegi Panjang: ", luas3.hitungLuasPPanjang(5,6))
		fmt.Println("--------------------------------")
	}
	// Soal 3
	{
		var samsung phone
		samsung.colors = samsung.coloring("Biru","Hitam")
		fmt.Println("Warna HP Samsung:", samsung.colors)

		var iPhone phone
		iPhone.colors = iPhone.coloring("Gold","Mint","Purple")
		fmt.Println("Warna HP iPhone:", iPhone.colors)
		fmt.Println("--------------------------------")
	}
	// Soal 4
	{
		var dataFilm = []movie{}

		tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
		tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
		tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
		tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)
		for i, item := range dataFilm {
				fmt.Println(i+1,".Judul :" ,item.title)
				fmt.Println("   Duration :" ,item.duration)
				fmt.Println("   Genre :" ,item.genre)
				fmt.Println("   Year :" ,item.year)
			}
	}
}