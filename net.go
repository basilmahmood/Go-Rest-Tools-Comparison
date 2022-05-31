package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func extractPathParam(pathParamName string, path string) (string, error) {
	var prev string

	for _, s := range strings.Split(path, "/") {
		if prev == pathParamName {
			return s, nil
		}
		prev = s
	}

	return "", errors.New("failed to extract path param")
}

func get(w http.ResponseWriter, r *http.Request) {
	url := r.URL
	repoName, err := extractPathParam(`repo`, url.Path)
	objectId, err := extractPathParam(`id`, url.Path)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(data[repoName+"/"+objectId])
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	var d interface{}

	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url := r.URL
	repoName, err := extractPathParam(`repo`, url.Path)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		data[repoName+"/"+strconv.Itoa(id)] = d

		w.Write([]byte(strconv.Itoa(id)))
		id++
	}
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		get(w, r)
	case "POST":
		post(w, r)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Only GET and POST supported"))
	}
}

func netAPI() {
	http.HandleFunc("/", routeHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
