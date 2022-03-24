package main

import (
	"Tugas14/config"
	"Tugas14/mahasiswa"
	"Tugas14/models"
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
	router.GET("/mahasiswa", GetMahasiswa)
	router.POST("/movie/create", PostMahasiswa)
	router.PUT("/movie/:id/update", UpdateMahasiswa)
	router.DELETE("/movie/:id/delete", DeleteMahasiswa)
  	fmt.Println("Server Running at Port 8080")
  	log.Fatal(http.ListenAndServe(":8080", router))
}

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
	var nilai models.NilaiMahasiswa
	if err := json.NewDecoder(r.Body).Decode(&nilai); err != nil {
	  utils.ResponseJSON(w, err, http.StatusBadRequest)
	  return
	}
	if err := mahasiswa.Insert(ctx, nilai); err != nil {
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
	var mhs models.NilaiMahasiswa
  
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