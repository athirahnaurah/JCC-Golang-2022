package main

import "fmt"

const PI = 3.14
// Soal 1
func hitungLuasLingkaran(r float64) *float64 {
	luas := PI * r * r
	return &luas
}

func hitungKelLingkaran(r float64) *float64{
	kel := 2 * PI * r
	return &kel
}

// Soal 2
func introduce(sentence *string, name, gender, job, age string){
	if (gender == "laki-laki"){
		gender = "Pak"
	}else{
		gender="Bu"
	}
	introduceText := gender + " " + name + " adalah seorang " + job + " yang berusia " + age + " tahun"
	*(sentence) = introduceText 
}

// Soal 3
func isiBuah (buah *[]string, namaBuah... string){
		*buah = append(*buah,namaBuah...)
}

// Soal 4
var tambahDataFilm = func (title, jam, genre, tahun string, data *[]map[string]string){
	dataFilm := map[string]string{
			"title": title,
			"jam": jam,
			"genre":  genre,
			"tahun" : tahun,
	}
	*data = append(*data,dataFilm)
}

func main() {
	// Soal 1
	{
		
		var luasLingkaran *float64 
		var kelilingLingkaran *float64
		
		var r float64 = 7.0
		kelilingLingkaran = hitungKelLingkaran(r)
		luasLingkaran = hitungLuasLingkaran(r)

		// tampilkan nilai

		fmt.Println("Luas (value) = ", *luasLingkaran)
		fmt.Println("Keliling (value) = ", *kelilingLingkaran)
		fmt.Println("--------------------------------")
		
	}
	// Soal 2
	{
		var sentence string 
		introduce(&sentence, "John", "laki-laki", "penulis", "30")

		fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
		introduce(&sentence, "Sarah", "perempuan", "model", "28")

		fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"
		fmt.Println("--------------------------------")
	}

	// Soal 3
	{
		var buah = []string{}
		isiBuah(&buah, "Jeruk","Semangka","Mangga","Strawberry","Durian","Manggis","Alpukat")
		for i, buah := range buah {
			fmt.Printf("%d. %s\n", i+1, buah)
		}
		fmt.Println("--------------------------------")
	}

	//Soal 4
	{
		var dataFilm = []map[string]string{}

		tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
		tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
		tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
		tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

		// loop over elements of slice
		for i, film := range dataFilm {
			// film is a map[string]
			fmt.Print(i+1, ".")
    		// loop over keys and values in the map
    		for key, value := range film {
        		fmt.Println(" ",key, " : ", value)
    		}
		}
		fmt.Println("--------------------------------")
	}
}

