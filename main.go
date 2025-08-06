package main

import (
	"encoding/json"
	"html/template"
	"log"
	"myapp/rps"
	"net/http"
	"strconv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index.html")
}

func PlayRound(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlaRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "   ")
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(out)

}

func main() {

	//handler
	http.HandleFunc("/play", PlayRound)
	http.HandleFunc("/", homePage)

	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func RenderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)

	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)

	if err != nil {
		log.Println(err)
		return
	}
}
