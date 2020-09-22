package handler

import (
	"strconv"
	"log"
	"path"
	"html/template"
	"net/http"
	"v2/entity"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Menampilkan Halaman HTML
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error bro, slow mo", http.StatusInternalServerError)
		return
	}

	// Data Dinamis HTML ( Passing Value )

	//data := entity.Product{ID: 1, Name: "Tesla", Price: 300000000, Stock: 7}

	// Slice of Struct

	data := []entity.Product{
		{ID: 1, Name: "Tesla", Price: 300000000, Stock: 11},
		{ID: 2, Name: "Jazz", Price: 250000000, Stock: 8},
		{ID: 3, Name: "Pajero", Price: 500000000, Stock: 1},
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "error lagi bro", http.StatusInternalServerError)
		return
	}
}

func GaHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Halo mamang"))
}

func BerHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Halo gaber"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)
	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}
	
	data := map[string]interface{}{
		"content": idNumb,
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error masqueh", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error masqueh", http.StatusInternalServerError)
		return
	}

}

func PostGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method

	switch method {
	case "GET":
		w.Write([]byte("ini adalah GET"))
	case "POST":
		w.Write([]byte("ini adalah POST"))
	default:
		http.Error(w, "Error 404 holyshit", http.StatusBadRequest)
	}
}

func Form(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views", "form.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error hadeuh", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error hadeuh", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "Error hadeuh", http.StatusBadRequest)
}

func Process(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Error hadeuh", http.StatusInternalServerError)
			return
		}
		name := r.Form.Get("name")
		message := r.Form.Get("message")

		data := map[string]interface{}{
			"name": name,
			"message": message,
		}

		tmpl, err := template.ParseFiles(path.Join("views", "result.html"), path.Join("views", "layout.html"))
		if err != nil {
			log.Println(err)
			http.Error(w, "Error hadeuh", http.StatusInternalServerError)
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error hadeuh", http.StatusInternalServerError)
			return
		}

		return
	}

	http.Error(w, "Error hadeuh", http.StatusBadRequest)
}