package main

import (
	"Tugas14/config"
	"Tugas14/mahasiswa"
	"Tugas14/matkul"
	"Tugas14/models"
	"Tugas14/nilai"
	"Tugas14/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
) 

func main() {  
	db, e := config.MySQL()

    if e != nil {
        log.Fatal(e)
    }

    eb := db.Ping()
    if eb != nil {
        panic(eb.Error()) 
    }

    fmt.Println("Success")
 	router := httprouter.New()

	router.GET("/matkul", GetMatkul)
	router.POST("/matkul/create", PostMatkul)
	router.PUT("/matkul/:id/update", UpdateMatkul)
	router.DELETE("/matkul/:id/delete", DeleteMatkul)

	router.GET("/mahasiswa", GetMahasiswa)
	router.POST("/mahasiswa/create", PostMahasiswa)
	router.PUT("/mahasiswa/:id/update", UpdateMahasiswa)
	router.DELETE("/mahasiswa/:id/delete", DeleteMahasiswa)

	router.GET("/nilai", GetNilai)
	router.POST("/nilai/create", PostNilai)
	router.PUT("/nilai/:id/update", UpdateNilai)
	router.DELETE("/nilai/:id/delete", DeleteNilai)

  	fmt.Println("Server Running at Port 8080")
  	log.Fatal(http.ListenAndServe(":8080", router))
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

// CRUD NILAI

// Read
func GetNilai(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	matkul, err := nilai.GetAll(ctx)
  
	if err != nil {
	  fmt.Println(err)
	}
  
	utils.ResponseJSON(w, matkul, http.StatusOK)
  }

  // Create
func PostNilai(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
		return
	  }
	  ctx, cancel := context.WithCancel(context.Background())
	  defer cancel()
	  var n models.Nilai
	  if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		utils.ResponseJSON(w, err, http.StatusBadRequest)
		return
	  }
	  if err := nilai.Insert(ctx, n); err != nil {
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	  }
	  res := map[string]string{
		"status": "Succesfully",
	  }
	  utils.ResponseJSON(w, res, http.StatusCreated)
}

  // Update
func UpdateNilai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
	  http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
	  return
	}
  
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var n models.Nilai
  
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
	  utils.ResponseJSON(w, err, http.StatusBadRequest)
	  return
	}
  
	var idNilai = ps.ByName("id")
  
	if err := nilai.Update(ctx, n, idNilai); err != nil {
	  utils.ResponseJSON(w, err, http.StatusInternalServerError)
	  return
	}
  
	res := map[string]string{
	  "status": "Succesfully",
	}
  
	utils.ResponseJSON(w, res, http.StatusCreated)
  }

  // Delete
func DeleteNilai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idNilai = ps.ByName("id")
	if err := mahasiswa.Delete(ctx, idNilai); err != nil {
	  kesalahan := map[string]string{
		"error": fmt.Sprintf("%v", err),
	  }
	  utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
	  return
	}
	res := map[string]string{
	  "status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
  }

// CRUD MAHASISWA 

// Read
func GetMahasiswa(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mahasiswa, err := mahasiswa.GetAll(ctx)
  
	if err != nil {
	  fmt.Println(err)
	}
  
	utils.ResponseJSON(w, mahasiswa, http.StatusOK)
  }

  // Create
func PostMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
	  http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
	  return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var mhs models.Mahasiswa
	if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
	  utils.ResponseJSON(w, err, http.StatusBadRequest)
	  return
	}
	if err := mahasiswa.Insert(ctx, mhs); err != nil {
	  utils.ResponseJSON(w, err, http.StatusInternalServerError)
	  return
	}
	res := map[string]string{
	  "status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
  }

  // Update
func UpdateMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
	  http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
	  return
	}
  
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var mhs models.Mahasiswa
  
	if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
	  utils.ResponseJSON(w, err, http.StatusBadRequest)
	  return
	}
  
	var idMhs = ps.ByName("id")
  
	if err := mahasiswa.Update(ctx, mhs, idMhs); err != nil {
	  utils.ResponseJSON(w, err, http.StatusInternalServerError)
	  return
	}
  
	res := map[string]string{
	  "status": "Succesfully",
	}
  
	utils.ResponseJSON(w, res, http.StatusCreated)
  }

  // Delete
func DeleteMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idMhs = ps.ByName("id")
	if err := mahasiswa.Delete(ctx, idMhs); err != nil {
	  kesalahan := map[string]string{
		"error": fmt.Sprintf("%v", err),
	  }
	  utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
	  return
	}
	res := map[string]string{
	  "status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
  }

// CRUD MATA KULIAH 

// Read
func GetMatkul(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	matkul, err := matkul.GetAll(ctx)
  
	if err != nil {
	  fmt.Println(err)
	}
  
	utils.ResponseJSON(w, matkul, http.StatusOK)
  }

  // Create
func PostMatkul(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
	  http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
	  return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var mk models.MataKuliah
	if err := json.NewDecoder(r.Body).Decode(&mk); err != nil {
	  utils.ResponseJSON(w, err, http.StatusBadRequest)
	  return
	}
	if err := matkul.Insert(ctx, mk); err != nil {
	  utils.ResponseJSON(w, err, http.StatusInternalServerError)
	  return
	}
	res := map[string]string{
	  "status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
  }

  // Update
func UpdateMatkul(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
	  http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
	  return
	}
  
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var mk models.MataKuliah
  
	if err := json.NewDecoder(r.Body).Decode(&mk); err != nil {
	  utils.ResponseJSON(w, err, http.StatusBadRequest)
	  return
	}
  
	var idMk = ps.ByName("id")
  
	if err := matkul.Update(ctx, mk, idMk); err != nil {
	  utils.ResponseJSON(w, err, http.StatusInternalServerError)
	  return
	}
  
	res := map[string]string{
	  "status": "Succesfully",
	}
  
	utils.ResponseJSON(w, res, http.StatusCreated)
  }

  // Delete
func DeleteMatkul(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idMk = ps.ByName("id")
	if err := mahasiswa.Delete(ctx, idMk); err != nil {
	  kesalahan := map[string]string{
		"error": fmt.Sprintf("%v", err),
	  }
	  utils.ResponseJSON(w, kesalahan, http.StatusInternalServerError)
	  return
	}
	res := map[string]string{
	  "status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusOK)
  }