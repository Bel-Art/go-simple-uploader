package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

type Filename struct {
	Name string `json:"name" required:"true"`
}

func respond(w http.ResponseWriter, status int, obj map[string]string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(obj)
}

func SendFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		respond(w, http.StatusBadRequest, map[string]string{"status": "ERROR_INVALID_METHOD"})
		return
	}
	var fileInfos Filename
	err := json.NewDecoder(r.Body).Decode(&fileInfos)
	if err != nil {
		respond(w, http.StatusBadRequest, map[string]string{"status": "ERROR_INVALID_BODY"})
		return
	}
	pathToFile := filepath.Join(".", "files", fileInfos.Name)
	fileBytes, err := ioutil.ReadFile(pathToFile)
	if err != nil {
		respond(w, http.StatusBadRequest, map[string]string{"status": "ERROR_READING_FILE"})
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
}

func ReceiveFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		respond(w, http.StatusBadRequest, map[string]string{"status": "ERROR_INVALID_METHOD"})
		return
	}
	r.ParseMultipartForm(32 << 20) // limit max input length
	var buf bytes.Buffer
	file, _, err := r.FormFile("file")
	if err != nil {
		respond(w, http.StatusBadRequest, map[string]string{"status": "ERROR_NO_FILE"})
		return
	}
	defer file.Close()
	io.Copy(&buf, file)
	id := uuid.New().String()
	pathToFile := filepath.Join(".", "files", id)
	f, err := os.Create(pathToFile)
	if err != nil {
		respond(w, http.StatusBadRequest, map[string]string{"status": "ERROR_FILE_CREATION"})
		return
	}
	f.Write(buf.Bytes())
	buf.Reset()
	respond(w, http.StatusOK, map[string]string{"status": "OK", "name": id})
}
