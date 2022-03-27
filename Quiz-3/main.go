package main

import (
	"Quiz-3/config"
	"Quiz-3/endpoints"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func BasicAuth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Get the Basic Authentication credentials
		user, password, hasAuth := r.BasicAuth()

		if hasAuth && user == "admin" && password == "password" {
			// Delegate request to the given handle
			h(w, r, ps)
		}else if hasAuth && user == "editor" && password == "secret" {
			h(w, r, ps)
		}else if hasAuth && user == "trainer" && password == "rahasia" {

		}else {
			// Request Basic Authentication otherwise
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}


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
	
	// Soal 3 & 5

	// API CATEGORY
	router.GET("/categories", endpoints.GetCategory)
	router.GET("/categories/:id", endpoints.GetBookFromCategory)
	router.POST("/categories", BasicAuth(endpoints.PostCategory))
	router.PUT("/categories/:id", BasicAuth(endpoints.UpdateCategory))
	router.DELETE("/categories/:id", BasicAuth(endpoints.DeleteCategory))
	
	// API BOOK
	router.GET("/books",endpoints.GetBook)
	router.POST("/books",BasicAuth(endpoints.PostBook))
	router.PUT("/books/:id", BasicAuth(endpoints.UpdateBook))
	router.DELETE("/books/:id", BasicAuth(endpoints.DeleteBook))

	// Soal 4
	router.GET("/book",endpoints.GetTitle)

  	fmt.Println("Server Running at Port 8080")
  	log.Fatal(http.ListenAndServe(":8080", router))
}