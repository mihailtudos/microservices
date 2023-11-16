package handlers

import (
	"fmt"
	"github.com/mihailtudos/product-images-api/files"
	"io"
	"log/slog"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
)

// Files is a handler for reading and writing files
type Files struct {
	log   slog.Logger
	store files.Storage
}

// NewFiles creates a new File handler
func NewFiles(s files.Storage, l slog.Logger) *Files {
	return &Files{store: s, log: l}
}

// UploadRest implements the http.Handler interface
func (f *Files) UploadRest(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.log.Info("Handle POST", "id", id, "filename", fn)

	// no need to check for invalid id or filename as the mux router will not send requests
	// here unless they have the correct parameters

	f.saveFile(id, fn, rw, r.Body)
}

// UploadMultiPart handler files received via multi-part forms data
func (f *Files) UploadMultiPart(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(128 * 1024)
	if err != nil {
		f.log.Error("Bad request", slog.String("error", err.Error()))
		http.Error(w, "Expected multipart form data", http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	_, idErr := strconv.Atoi(vars["id"])
	if idErr != nil {
		f.log.Error("Bad request", slog.String("error", err.Error()))
		http.Error(w, "Expected integer id", http.StatusBadRequest)
		return
	}

	file, fh, err := r.FormFile("file")
	if err != nil {
		f.log.Error("Bad request", slog.String("error", err.Error()))
		http.Error(w, "Expected file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileType := filepath.Ext(fh.Filename)
	fileName := fmt.Sprintf("main%s", fileType)

	f.saveFile(vars["id"], fileName, w, file)
	f.log.Info("processed form", slog.String("id", vars["id"]), slog.String("file_name", fileName))
}

func (f *Files) invalidURI(uri string, rw http.ResponseWriter) {
	f.log.Error("Invalid path", "path", uri)
	http.Error(rw, "Invalid file path should be in the format: /[id]/[filepath]", http.StatusBadRequest)
}

// saveFile saves the contents of the request to a file
func (f *Files) saveFile(id, path string, rw http.ResponseWriter, r io.ReadCloser) {
	f.log.Info("Save file for product", "id", id, "path", path)

	fp := filepath.Join(id, path)

	err := f.store.Save(fp, r)
	if err != nil {
		f.log.Error("Unable to save file", "error", err)
		http.Error(rw, "Unable to save file", http.StatusInternalServerError)
	}
}
