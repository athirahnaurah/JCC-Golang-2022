package endpoints

import (
	"Quiz-3/book"
	"Quiz-3/category"
	"Quiz-3/models"
	"Quiz-3/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
)

// Read
func GetCategory(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	category, err := category.GetAll(ctx)
  
	if err != nil {
	  fmt.Println(err)
	}
  
	utils.ResponseJSON(w, category, http.StatusOK)
  }

  // Create
func PostCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
	  http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
	  return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var _category models.Category
	if err := json.NewDecoder(r.Body).Decode(&_category); err != nil {
	  utils.ResponseJSON(w, err, http.StatusBadRequest)
	  return
	}

	if err := category.Insert(ctx, _category); err != nil {
	  utils.ResponseJSON(w, err, http.StatusInternalServerError)
	  return
	}
	res := map[string]string{
	  "status": "Succesfully",
	}
	utils.ResponseJSON(w, res, http.StatusCreated)
  }

  // Update
func UpdateCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
	  http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
	  return
	}
  
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var _category models.Category
  
	if err := json.NewDecoder(r.Body).Decode(&_category); err != nil {
	  utils.ResponseJSON(w, err, http.StatusBadRequest)
	  return
	}
  
	var idCategory = ps.ByName("id")
  
	if err := category.Update(ctx, _category, idCategory); err != nil {
	  utils.ResponseJSON(w, err, http.StatusInternalServerError)
	  return
	}
  
	res := map[string]string{
	  "status": "Succesfully",
	}
  
	utils.ResponseJSON(w, res, http.StatusCreated)
  }

  // Delete
func DeleteCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idCategory = ps.ByName("id")
	if err := category.Delete(ctx, idCategory); err != nil {
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

// Read
func GetBook(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	book, err := book.GetAll(ctx)
  
	if err != nil {
	  fmt.Println(err)
	}
  
	utils.ResponseJSON(w, book, http.StatusOK)
  }

  // Read
func GetBookFromCategory(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idBook = ps.ByName("id")
	book, err := book.GetBookFromCategory(ctx,idBook)
  
	if err != nil {
	  fmt.Println(err)
	}
  
	utils.ResponseJSON(w, book, http.StatusOK)
  }

  // Create
func PostBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
	  http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
	  return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var _book models.Book
	if err := json.NewDecoder(r.Body).Decode(&_book); err != nil {
	  utils.ResponseJSON(w, err, http.StatusBadRequest)
	  return
	}
	_, err := url.ParseRequestURI(_book.Image_url)
   	
	if _book.Release_year < 1980 || _book.Release_year > 2021 || err!= nil{
		err := map[string]string{
			"status": "Failed",
			"message" : "release year is not allowed",
		  }
		utils.ResponseJSON(w, err, http.StatusInternalServerError)
		return
	}else{
		_book.Thickness = checkThickness(_book.Total_page)

		if err := book.Insert(ctx, _book); err != nil {
		  utils.ResponseJSON(w, err, http.StatusInternalServerError)
		  return
		}
		res := map[string]string{
		  "status": "Succesfully",
		}
		utils.ResponseJSON(w, res, http.StatusCreated)
	}
  }

func checkThickness(page int) string{
	var thickness string
	if page <= 100 {
		thickness = "tipis"
	} else if page >= 101 && page <=200 {
		thickness = "sedang"
	} else if page >= 201 {
		thickness = "tebal"
	}
	return thickness
}

  // Update
func UpdateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
	  http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
	  return
	}
  
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var _book models.Book
  
	if err := json.NewDecoder(r.Body).Decode(&_book); err != nil {
	  utils.ResponseJSON(w, err, http.StatusBadRequest)
	  return
	}
  
	var idBook = ps.ByName("id")
  
	if err := book.Update(ctx, _book, idBook); err != nil {
	  utils.ResponseJSON(w, err, http.StatusInternalServerError)
	  return
	}
  
	res := map[string]string{
	  "status": "Succesfully",
	}
  
	utils.ResponseJSON(w, res, http.StatusCreated)
  }

  // Delete
func DeleteBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var idBook= ps.ByName("id")
	if err := book.Delete(ctx, idBook); err != nil {
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

func GetTitle(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var title = r.URL.Query().Get("title")
	book, err := book.GetTitle(ctx,title)
  
	if err != nil {
	  fmt.Println(err)
	}
  
	utils.ResponseJSON(w, book, http.StatusOK)

}