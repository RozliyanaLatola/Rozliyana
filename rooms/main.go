package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

// Ruang struct (Model) ...
type Ruang struct {
	RuangID    	 string `json:"RuangID"`
	RuangName    string `json:"RuangName"`
	PasienName   string `json:"PasienName"`
	Biaya      	 string `json:"Biaya"`
	LamaNginap	 string `json:"LamaNginap"`
	Penyakit  	 string `json:"Penyakit"`
}

// Get all ruang

func getRuangs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var ruangs []Ruang

	sql := `SELECT
				RuangID,
				IFNULL(RuangName,''),
				IFNULL(PasienName,'') PasienName,
				IFNULL(Biaya,'') Biaya,
				IFNULL(LamaNginap,'') LamaNginap,
				IFNULL(Penyakit,'') Penyakit
			FROM ruangs`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		var ruang Ruang
		err := result.Scan(&ruang.RuangID, &ruang.RuangName, &ruang.PasienName,
			&ruang.Biaya, &ruang.LamaNginap, &ruang.Penyakit)

		if err != nil {
			panic(err.Error())
		}
		ruangs = append(ruangs, ruang)
	}

	json.NewEncoder(w).Encode(ruangs)
}

func createRuang(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		RuangID := r.FormValue("RuangID")
		RuangName := r.FormValue("RuangName")
		PasienName := r.FormValue("PasienName")
		Biaya := r.FormValue("Biaya")
		LamaNginap := r.FormValue("LamaNginap")
		Penyakit := r.FormValue("Penyakit")

		stmt, err := db.Prepare("INSERT INTO ruangs (RuangID,RuangName,PasienName,Biaya,LamaNginap,Penyakit) VALUES (?,?,?,?,?,?)")

		_, err = stmt.Exec(RuangID, RuangName, PasienName, Biaya, LamaNginap, Penyakit)

		if err != nil {
			fmt.Fprintf(w, "Data Duplicate")
		} else {
			fmt.Fprintf(w, "Data Created")
		}

	}
}

func getRuang(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ruangs []Ruang
	params := mux.Vars(r)

	sql := `SELECT
				RuangID,
				IFNULL(RuangName,''),
				IFNULL(PasienName,'') PasienName,
				IFNULL(Biaya,'') Biaya,
				IFNULL(LamaNginap,'') LamaNginap,
				IFNULL(Penyakit,'') Penyakit
			FROM ruangs WHERE RuangID = ?`

	result, err := db.Query(sql, params["id"])

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var ruang Ruang

	for result.Next() {

		err := result.Scan(&ruang.RuangID, &ruang.RuangName, &ruang.PasienName,
			&ruang.Biaya, &ruang.LamaNginap, &ruang.Penyakit)

		if err != nil {
			panic(err.Error())
		}

		ruangs = append(ruangs, ruang)
	}

	json.NewEncoder(w).Encode(ruangs)
}

func updateRuang(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {

		params := mux.Vars(r)

		newRuangName := r.FormValue("RuangName")
		newPasienName := r.FormValue("PasienName")
		newBiaya := r.FormValue("Biaya")
		newLamaNginap := r.FormValue("LamaNginap")
		newPenyakit := r.FormValue("Penyakit")

		stmt, err := db.Prepare("UPDATE ruangs SET RuangName = ?, PasienName = ?, Biaya = ?, LamaNginap = ?, Penyakit = ? WHERE RuangID = ?")

		_, err = stmt.Exec(newRuangName, newPasienName, newBiaya, newLamaNginap, newPenyakit, params["id"])

		if err != nil {
			fmt.Fprintf(w, "Data not found or Request error")
		}

		fmt.Fprintf(w, "Ruang with RuangID = %s was updated", params["id"])
	}
}

func delRuang(w http.ResponseWriter, r *http.Request) {

	RuangID := r.FormValue("RuangID")
	RuangName := r.FormValue("RuangName")

	stmt, err := db.Prepare("DELETE FROM ruangs WHERE RuangID = ? AND RuangName = ?")

	_, err = stmt.Exec(RuangID, RuangName)

	if err != nil {
		fmt.Fprintf(w, "delete failed")
	}

	fmt.Fprintf(w, "Ruang with ID = %s was deleted", RuangID)
}

func getPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var ruangs []Ruang

	RuangID := r.FormValue("RuangID")
	RuangName := r.FormValue("RuangName")

	sql := `SELECT
				RuangID,
				IFNULL(RuangName,''),
				IFNULL(RuangName,'') RuangName,
				IFNULL(Biaya,'') Biaya,
				IFNULL(LamaNginap,'') LamaNginap,
				IFNULL(Penyakit,'') Penyakit
			FROM ruangs WHERE RuangID = ? AND RuangName = ?`

	result, err := db.Query(sql, RuangID, RuangName)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var ruang Ruang

	for result.Next() {

		err := result.Scan(&ruang.RuangID, &ruang.RuangName, &ruang.PasienName,
			&ruang.Biaya, &ruang.LamaNginap, &ruang.Penyakit)

		if err != nil {
			panic(err.Error())
		}

		ruangs = append(ruangs, ruang)
	}

	json.NewEncoder(w).Encode(ruangs)

}

// Main function
func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lyana")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// Init router
	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/ruangs", getRuangs).Methods("GET")
	r.HandleFunc("/ruangs/{id}", getRuang).Methods("GET")
	r.HandleFunc("/ruangs", createRuang).Methods("POST")
	r.HandleFunc("/ruangs/{id}", updateRuang).Methods("PUT")
	r.HandleFunc("/delruangs", delRuang).Methods("POST")

	//New
	r.HandleFunc("/getruang", getPost).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", r))
}
