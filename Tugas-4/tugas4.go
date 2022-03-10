package main

import "fmt"

func main() {
	// Soal 1
	{
		for angka := 1; angka <= 20; angka++ {
			if angka%2 != 0 && angka%3 == 0  {
				fmt.Print(angka)
				fmt.Println(" - I Love Coding")
			}else if angka%2 == 0{
				fmt.Print(angka)
				fmt.Println(" - Candradimuka")
			}else if angka%2!=0{
				fmt.Print(angka)
				fmt.Println(" - JCC")
			}
		}
		fmt.Println("------------------------")
	}

	// Soal 2
	{
		for i := 1; i <= 7; i++{
			for j := 1; j <= i; j++{
				fmt.Print("#")
			}
			fmt.Println() 
		}
		fmt.Println("------------------------")
	}

	// Soal 3
	{
		var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
		fmt.Println(kalimat[2:7])
		fmt.Println("------------------------")
	}
	// Soal 4
	{
		var sayuran = []string{
			"Bayam",
			"Buncis",
			"Kangkung",
			"Kubis",
			"Seledri",
			"Tauge",
			"Timun",
		}
		for i, sayuran := range sayuran {
			fmt.Printf("%d.%s\n", i+1, sayuran)
		}
		fmt.Println("------------------------")
	}

	// Soal 5
	{
		var vol = 1
		var satuan = map[string]int{
			"panjang": 7,
			"lebar":   4,
			"tinggi":  6,
		  }
		for i, satuan := range satuan {
			fmt.Println(i," = ", satuan)
			vol *= satuan
		}
		fmt.Print("Volume Balok = ", vol)
	}
}