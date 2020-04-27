package serve

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func writeJson(w http.ResponseWriter, obj interface{}) {
	jsn, err := json.Marshal(obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(jsn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func putJson(w http.ResponseWriter, r io.ReadCloser, obj interface{}) {
	defer func() {
		err := r.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()

	buf := bytes.Buffer{}
	_, err := buf.ReadFrom(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(buf.Bytes(), obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
