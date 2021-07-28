package main

import (
	"encoding/json"
	"net/http"
)

type Curso struct {
	Title        string `json:"title"`
	NumeroVideos int    `json:"numero_videos"`
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		cursos := Cursos{
			Curso{"Curso de GO", 30},
			Curso{"Curso de JS", 20},
			Curso{"Curso de PY", 25},
		}
		json.NewEncoder(w).Encode(cursos)
	})
	http.ListenAndServe(":8000", nil)
}
