package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	temp, err := template.ParseGlob("*.html")
	if err != nil {
		fmt.Println(fmt.Printf("ERREUR => %s", err.Error()))
		return
	}

	type Etudiant struct {
		Prenom string
		Nom    string
		Age    int
		Genre  string
	}

	type Variable struct {
		NomPromotion     string
		FilierePromotion string
		NiveauPromotion  int
		NombreEtudiants  int
		Etudiants        []Etudiant
	}

	type Number struct {
		Odd int
		Num int
	}

	type Form struct {
		Name     string
		Nickname string
		Date     string
		Genre    string
	}

	num := Number{0, 0}

	EtudiantsPromo := []Etudiant{{"RODRIGUES", "Cyril", 22, "Homme"}, {"MEDERREG", "Kheir-eddine", 22, "Homme"}, {"PHILIPIERT", "Alan", 26, "Homme"}}

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		data := Variable{"Mentor'ac",
			"Informatique",
			5,
			3,
			EtudiantsPromo}
		temp.ExecuteTemplate(w, "promo", data)
	})

	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {
		num.Num += 1
		num.Odd = num.Num % 2
		temp.ExecuteTemplate(w, "change", num)
	})

	http.HandleFunc("/user/init", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "login", 0)
	})

	http.HandleFunc("/user/redirection", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "redirection", 0)
	})

	http.HandleFunc("/form/post", func(rw http.ResponseWriter, r *http.Request) {
		infos := Form{r.FormValue("prenom"), r.FormValue("nom"), r.FormValue("date"), r.FormValue("genre")}
		temp.ExecuteTemplate(rw, "reponse", infos)
	})

	http.HandleFunc("/user/display", func(w http.ResponseWriter, r *http.Request) {
		infos := r
		temp.ExecuteTemplate(w, "display", infos)
	})

	fileserver := http.FileServer(http.Dir("./asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.ListenAndServe("localhost:6969", nil)
}
