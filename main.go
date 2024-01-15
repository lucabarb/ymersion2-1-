package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	temp, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		fmt.Printf(fmt.Sprintf("ERREUR => %s", err.Error()))
		return
	}

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "home", nil)
	})

	http.HandleFunc("/code/treatment", func(w http.ResponseWriter, r *http.Request) {
		// Utilisez r.FormValue("input") pour récupérer la valeur du formulaire
		input := r.FormValue("input")
		result := decode(input)
		fmt.Fprintf(w, "Résultat : %s", result)
	})

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	fmt.Println("Serveur démarré sur le port 8080...")
	http.ListenAndServe("localhost:8080", nil)
}

func decode(input string) string {
	decode := make([]byte, len(input))

	for i := 0; i < len(input); i++ {
		char := input[i]
		decodeChar := char

		if char >= 'A' && char <= 'Z' {
			decodeChar = (char-'A'-2+26)%26 + 'A'
		}

		if char >= 'a' && char <= 'z' {
			decodeChar = (char-'a'-2+26)%26 + 'a'
		}

		decode[i] = decodeChar
	}

	return string(decode)
}
