package controllers

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

func GetById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
	}
	idFromURL := r.URL.Query().Get("id")
	if idFromURL == "" {
		http.Error(w, "No data entered", http.StatusBadRequest)
	}
	id := strings.Split(idFromURL, ",")
	m := Reader()
	resp := [][]string{}
	for _, i := range id {
		a := m[i]
		resp = append(resp, a)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func Reader() map[string][]string {
	m := make(map[string][]string)
	file, err := os.Open("file.cvs")
	if err != nil {
		log.Fatalf("failed to open cvs file %v", err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	csvReader.FieldsPerRecord = -1
	csvReader.Comment = '#'
	csvReader.LazyQuotes = true
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatalf("failed to read cvs file %v", err)
	}
	for i := range data {
		m[data[i][0]] = data[i]
	}
	return m
}
