package main

import "fmt"
func main(){
	// Soal 1
	panjang := 12
	lebar := 4
	tinggi := 8
	  
	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas) 
	fmt.Println(keliling)
	fmt.Println(volume)

	// Soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// Soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("john", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"

	// Soal 4
	var dataFilm = []map[string]string{}
	// buatlah closure function disini
	var tambahDataFilm = func (title, jam, genre, tahun string){
		data := []map[string]string{
			{
				"title": title,
				"jam": jam,
				"genre":  genre,
				"tahun" : tahun,
			},
		}
		dataFilm = append(dataFilm,data...)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
	fmt.Println(item)
	}
}

// Soal 1
func luasPersegiPanjang(p int, l int) int {
	luas := p * l
	return luas
}

func kelilingPersegiPanjang(p int, l int) int {
	keliling := (2*p) + (2*l)
	return keliling
}

func volumeBalok(p int, l int, t int) int {
	vol := p * l * t
	return vol
}

// Soal 2
func introduce (name, gender, job, age string ) (result string){
	if gender == "laki-laki"{
		gender = "Pak"
	}else if gender == "perempuan"{
		gender ="Bu"
	}
	result = gender + " " + name + " adalah seorang " + job + " yang berusia " + age + " tahun"
	return result
}

// Soal 3
func buahFavorit(name string, buah...string) (result string){
	var buahfav string
	for _, buahstr := range buah {
		buahfav += `"` + buahstr + `"` + ", "
	}
	//remove "," in last char
	if last := len(buahfav) - 2; last >= 0 && buahfav[last] == ',' {
        buahfav = buahfav[:last]
    }
	result = "halo nama saya " + name +" dan buah favorit saya adalah " + buahfav
	return result
}