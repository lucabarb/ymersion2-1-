package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
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
	})

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	fmt.Println("Serveur démarré sur le port 8080...")
	http.ListenAndServe("localhost:8080", nil)
}
func addTwo(str string) string {
	var res strings.Builder
	for _, c := range str {
		newChar := c + 2
		if newChar > 'z' {
			newChar = newChar - 26
		}
		res.WriteRune(newChar)
	}
	return res.String()
}
