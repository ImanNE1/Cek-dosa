package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Hasil struct {
	TampilkanHasil bool
	Nama           string
	Skor           int
	Pesan          string
}

var rRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func cekHandler(w http.ResponseWriter, r *http.Request) {
	tmplt, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Error: Tidak bisa memuat template.", http.StatusInternalServerError)
		return
	}

	data := Hasil{TampilkanHasil: false}

	if r.Method == http.MethodPost {
		nama := r.FormValue("nama")

		if nama == "" {
			nama = "Seseorang"
		}
		skor := rRand.Intn(1000000) + 1
		var pesan string
		if skor <= 200000 {
			pesan = "Level dosamu masih sangat rendah. Aman!"
		} else if skor <= 500000 {
			pesan = "Sudah mulai terkumpul, saatnya lebih berhati-hati."
		} else if skor <= 800000 {
			pesan = "Level yang cukup mengkhawatirkan. Sering-sering introspeksi, ya!"
		} else {
			pesan = "Wow! Level dosamu sudah setara Sultan!"
		}

		data = Hasil{
			TampilkanHasil: true,
			Nama:           nama,
			Skor:           skor,
			Pesan:          pesan,
		}
	}
	tmplt.Execute(w, data)
}

func main() {
	http.HandleFunc("/", cekHandler)

	fmt.Println("Server 'Cek Dosa' berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
