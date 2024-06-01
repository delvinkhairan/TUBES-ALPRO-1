package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strconv"
)

type Negara struct {
	Nama           string
	JumlahEmas     int
	JumlahPerak    int
	JumlahPerunggu int
}

type Peringkat struct {
	NamaNegara     string
	JumlahEmas     int
	JumlahPerak    int
	JumlahPerunggu int
	TotalMedali    int
}

var daftarNegara [11]Negara
var peringkat [11]Peringkat

var jumlahNegara int

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/edit", editHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/ranking", rankingHandler)
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/sort", sortHandler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, daftarNegara[:jumlahNegara])
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		nama := r.FormValue("nama")
		emas, _ := strconv.Atoi(r.FormValue("emas"))
		perak, _ := strconv.Atoi(r.FormValue("perak"))
		perunggu, _ := strconv.Atoi(r.FormValue("perunggu"))
		jumlahNegara = tambahNegara(nama, emas, perak, perunggu, jumlahNegara)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.ServeFile(w, r, "add.html")
	}
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		nama := r.FormValue("nama")
		emas, _ := strconv.Atoi(r.FormValue("emas"))
		perak, _ := strconv.Atoi(r.FormValue("perak"))
		perunggu, _ := strconv.Atoi(r.FormValue("perunggu"))
		ubahNegara(nama, emas, perak, perunggu, jumlahNegara)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.ServeFile(w, r, "edit.html")
	}
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		nama := r.FormValue("nama")
		jumlahNegara = hapusNegara(nama, jumlahNegara)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.ServeFile(w, r, "delete.html")
	}
}

func rankingHandler(w http.ResponseWriter, r *http.Request) {
	// Sort countries by gold, silver, and bronze
	sort.Slice(daftarNegara[:jumlahNegara], func(i, j int) bool {
		if daftarNegara[i].JumlahEmas != daftarNegara[j].JumlahEmas {
			return daftarNegara[i].JumlahEmas > daftarNegara[j].JumlahEmas
		}
		if daftarNegara[i].JumlahPerak != daftarNegara[j].JumlahPerak {
			return daftarNegara[i].JumlahPerak > daftarNegara[j].JumlahPerak
		}
		return daftarNegara[i].JumlahPerunggu > daftarNegara[j].JumlahPerunggu
	})

	tmpl := template.Must(template.New("ranking.html").Funcs(template.FuncMap{
		"inc": inc,
	}).ParseFiles("ranking.html"))

	err := tmpl.Execute(w, daftarNegara[:jumlahNegara])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		rank, _ := strconv.Atoi(r.FormValue("rank"))
		temukanNegaraDariPeringkat(rank, jumlahNegara)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.ServeFile(w, r, "search.html")
	}
}

func sortHandler(w http.ResponseWriter, r *http.Request) {
	urutkanNegaraAlphabet(jumlahNegara)
	http.Redirect(w, r, "/", http.StatusFound)
}

func tambahNegara(nama string, emas, perak, perunggu, jumlahNegara int) int {
	if jumlahNegara < len(daftarNegara) {
		daftarNegara[jumlahNegara] = Negara{Nama: nama, JumlahEmas: emas, JumlahPerak: perak, JumlahPerunggu: perunggu}
		jumlahNegara++
		fmt.Println("Negara ditambahkan.")
	} else {
		fmt.Println("Kapasitas negara penuh.")
	}
	return jumlahNegara
}

func ubahNegara(nama string, emas, perak, perunggu, jumlahNegara int) {
	for i := 0; i < jumlahNegara; i++ {
		if daftarNegara[i].Nama == nama {
			daftarNegara[i].JumlahEmas = emas
			daftarNegara[i].JumlahPerak = perak
			daftarNegara[i].JumlahPerunggu = perunggu
			fmt.Println("Data Negara diubah.")
			return
		}
	}
	fmt.Println("Negara tidak ditemukan.")
}

func hapusNegara(nama string, jumlahNegara int) int {
	for i := 0; i < jumlahNegara; i++ {
		if daftarNegara[i].Nama == nama {
			for j := i; j < jumlahNegara-1; j++ {
				daftarNegara[j] = daftarNegara[j+1]
			}
			jumlahNegara--
			fmt.Println("Negara dihapus.")
			return jumlahNegara
		}
	}
	fmt.Println("Negara tidak ditemukan.")
	return jumlahNegara
}

func updatePeringkat(jumlahNegara int) {
	for i := 0; i < jumlahNegara; i++ {
		totalMedali := daftarNegara[i].JumlahEmas + daftarNegara[i].JumlahPerak + daftarNegara[i].JumlahPerunggu
		peringkat[i] = Peringkat{
			NamaNegara:     daftarNegara[i].Nama,
			JumlahEmas:     daftarNegara[i].JumlahEmas,
			JumlahPerak:    daftarNegara[i].JumlahPerak,
			JumlahPerunggu: daftarNegara[i].JumlahPerunggu,
			TotalMedali:    totalMedali,
		}
	}
	for i := 0; i < jumlahNegara; i++ {
		for j := i + 1; j < jumlahNegara; j++ {
			if peringkat[j].TotalMedali > peringkat[i].TotalMedali {
				peringkat[i], peringkat[j] = peringkat[j], peringkat[i]
			}
		}
	}
}

func urutkanNegaraAlphabet(jumlahNegara int) {
	for i := 1; i < jumlahNegara; i++ {
		idx_min := daftarNegara[i]
		j := i - 1
		for j >= 0 && daftarNegara[j].Nama > idx_min.Nama {
			daftarNegara[j+1] = daftarNegara[j]
			j--
		}
		daftarNegara[j+1] = idx_min
	}
}

func temukanNegaraDariPeringkat(rank, jumlahNegara int) {
	if rank > 0 && rank <= jumlahNegara {
		fmt.Println("Negara pada peringkat", rank, "adalah", peringkat[rank-1].NamaNegara)
	} else {
		fmt.Println("Peringkat tidak valid.")
	}
}

func inc(i int) int {
	return i + 1
}
