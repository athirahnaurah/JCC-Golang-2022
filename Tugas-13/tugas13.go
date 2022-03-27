package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type NilaiMahasiswa struct {
	Nama string `json:"nama"`
	MataKuliah string `json:"matkul"`
	IndeksNilai string `json:"indeks"`
	Nilai uint64 `json:"nilai"`
	ID uint64 `json:"id"`
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	  uname, pwd, ok := r.BasicAuth()
	  if !ok {
		w.Write([]byte("Username atau Password tidak boleh kosong"))
		return
	  }
  
	  if uname == "admin" && pwd == "admin" {
		next.ServeHTTP(w, r)
		return
	 }
	  w.Write([]byte("Username atau Password tidak sesuai"))
	  return
	})
  }

func getMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
	  mhs := nilaiNilaiMahasiswa
	  dataMahasiswa, err := json.Marshal(mhs)
	  
	  if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	  }
  
	  w.Header().Set("Content-Type", "application/json")
	  w.WriteHeader(http.StatusOK)
	  w.Write(dataMahasiswa)
	  return
	}
	http.Error(w, "ERROR....", http.StatusNotFound)
  }

func PostMahasiswa(w http.ResponseWriter, r *http.Request) {
	var Mhs NilaiMahasiswa
	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&Mhs); err != nil {
				log.Fatal(err)
			  }
	}else{
		idMhs := len(nilaiNilaiMahasiswa)+1
		var id uint64 = uint64(idMhs)
		nama := r.PostFormValue("nama")
		matkul := r.PostFormValue("matkul")
		getNilai := r.PostFormValue("nilai")
		nilai, _ := strconv.ParseUint(getNilai, 10, 64)
		
		indeks := checkIndeks(uint(nilai))
		if nilai <= 100 {
			Mhs = NilaiMahasiswa{
				ID: id,
				Nama : nama,
				MataKuliah: matkul,
				Nilai: nilai,
				IndeksNilai: indeks,
			}
			nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa,Mhs)
		}else{
			w.Write([]byte("Nilai tidak boleh lebih dari 100"))
		}
		

		dataMhs, _ := json.Marshal(nilaiNilaiMahasiswa) 
    	w.Write(dataMhs)                
    	return
	}
}
}

func checkIndeks( nilai uint)string{
	var indeks string
	if nilai >= 80 {
		indeks = "A"
	}else if nilai >= 70 && nilai < 80{
		indeks = "B"
	}else if nilai >= 60 && nilai < 70{
		indeks = "C"
	}else if nilai >= 50 && nilai < 60{
		indeks = "D"
	}else if nilai < 50{
		indeks = "E"
	}
	return indeks
}


func main() {
	router := httprouter.New()
	
	http.Handle("/post_mhs", Auth(http.HandlerFunc(PostMahasiswa)))
	http.HandleFunc("/get_mhs", getMahasiswa)
	router.GET("",getMahasiswa)
	fmt.Println("server running at http://localhost:8080")

  if err := http.ListenAndServe(":8080", router); err != nil {
    log.Fatal(err)
  }

}