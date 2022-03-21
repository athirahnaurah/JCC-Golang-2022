package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// Soal 1
func displayPhone(phones []string, second int, wg *sync.WaitGroup){
	for i, phone := range phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(time.Second*time.Duration(second))
	}
	wg.Done()
}

// Soal 2
func getMovies(moviesChannel chan <- string, movies... string){
	ch := make(chan int, 6)
	fmt.Println("List Movies:")
	for i , movie := range movies{
		ch <- i+1
		fmt.Print(<-ch,". ")
		moviesChannel <- movie
		time.Sleep(time.Microsecond)
	}
	close(moviesChannel)
	
}

// Soal 3
func luasLingkaran(ch1 chan <- float64, r float64){
	 ch1 <- r * r * math.Pi
}

func kelilingLingkaran(ch2 chan <- float64, r float64){
	ch2 <- 2 * math.Pi * r
}

func volumeTabung(ch3 chan <- float64, r float64, t float64){
	ch3 <- math.Pi * math.Pow(r,2) * t
}

func luasPersegiPanjang(ch chan <- int, p int, l int){
	ch <- p * l
}

func kelilingPersegiPanjang(ch chan <- int, p int, l int){
	ch <- (2 * p) + (2 * l)
}

func volumeBalok(ch chan <- int, p int, l int, t int){
	ch <- p * l * t
}

func main(){
	// Soal 1
	{
		var wg sync.WaitGroup
		var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
		wg.Add(1)
		go displayPhone(phones,1,&wg)
		wg.Wait()

	}
	// Soal 2
	{
		var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

		moviesChannel := make(chan string)

		go getMovies(moviesChannel, movies...)
		
		for value := range moviesChannel {
			fmt.Println(value)
		}
	}
	// Soal 3
	{
		ch:= make(chan float64,3)
		go luasLingkaran(ch,8)
		fmt.Println("Luas: ",<-ch)

		go kelilingLingkaran(ch,14)
		fmt.Println("Keliling: ",<-ch)

		go volumeTabung(ch,20,10)
		fmt.Println("Volume: ",<-ch)
	}
	// Soal 4
	{
		ch1:=make(chan int)
		go luasPersegiPanjang(ch1,10,5)

		ch2:=make(chan int)
		go kelilingPersegiPanjang(ch2,10,5)

		ch3:=make(chan int)
		go volumeBalok(ch3,10,5,3)

		for i := 0; i < 3; i++ {
			select {
			case luas := <-ch1:
			  	fmt.Printf("Luas \t: %d \n", luas)
			case keliling := <-ch2:
			  	fmt.Printf("Keliling: %d \n", keliling)
			case volume := <- ch3:
				fmt.Printf("Volume \t: %d \n", volume)
			}
		  }
	}
}