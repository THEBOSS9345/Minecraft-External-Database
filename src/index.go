package main

import (
	database "Minecraft-External-Database/src/Database"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func main() {
	log.Println("Starting server on :7240")
	http.HandleFunc("/set", setHandler)
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/has", hasHandler)
	http.HandleFunc("/list", listHandler)
	log.Println("Server started on :7240")
	log.Fatal(http.ListenAndServe(":7240", nil))
}

func getDatabaseFile(r *http.Request) (string, error) {
	file := r.URL.Query().Get("file")
	if file == "" {
		return "", http.ErrMissingFile
	}
	return file, nil
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	dbFile, err := getDatabaseFile(r)
	if err != nil {
		http.Error(w, "file parameter is required", http.StatusBadRequest)
		return
	}

	db, err := database.New(dbFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var data map[string]interface{}
	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for key, value := range data {
		err = db.Set(key, value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	dbFile, err := getDatabaseFile(r)
	if err != nil {
		http.Error(w, "file parameter is required", http.StatusBadRequest)
		return
	}

	db, err := database.New(dbFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "key is required", http.StatusBadRequest)
		return
	}

	value, err := db.Get(key)
	if err != nil {

		if strings.Contains(err.Error(), "sql: no rows in result set") {
			http.Error(w, "key not found", http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(value)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	dbFile, err := getDatabaseFile(r)
	if err != nil {
		http.Error(w, "file parameter is required", http.StatusBadRequest)
		return
	}

	db, err := database.New(dbFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "key is required", http.StatusBadRequest)
		return
	}

	err = db.Delete(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func hasHandler(w http.ResponseWriter, r *http.Request) {
	dbFile, err := getDatabaseFile(r)
	if err != nil {
		http.Error(w, "file parameter is required", http.StatusBadRequest)
		return
	}

	db, err := database.New(dbFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "key is required", http.StatusBadRequest)
		return
	}

	exists, err := db.Has(key)
	if err != nil {

		if strings.Contains(err.Error(), "sql: no rows in result set") {
			json.NewEncoder(w).Encode(map[string]bool{"exists": false})
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]bool{"exists": exists})
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	dbFile, err := getDatabaseFile(r)
	if err != nil {
		http.Error(w, "file parameter is required", http.StatusBadRequest)
		return
	}

	db, err := database.New(dbFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	keys, err := db.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(keys)
}
