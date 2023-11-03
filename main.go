package main

import (
	"encoding/json"
	"github.com/h2non/filetype"
	"github.com/h2non/filetype/matchers"
	"gocv.io/x/gocv"
	"io"
	"log"
	"net/http"
)

func main() {
	acceptFileType := matchers.Map{
		matchers.TypeJpeg: matchers.Jpeg,
		matchers.TypePng:  matchers.Png,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := r.FormValue("url")
		var f io.Reader
		if url == "" {
			var err error
			f, _, err = r.FormFile("file")
			for k := range r.Form {
				log.Print(k)
			}
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		} else {
			resp, err := http.Get(url)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			f = resp.Body
		}
		b, err := io.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if !filetype.MatchesMap(b, acceptFileType) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		m, err := gocv.IMDecode(b, gocv.IMReadColor)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer m.Close()
		json.NewEncoder(w).Encode(WxDecodeQrCodeMulti(m))
	})

	log.Print("server start @ :8080")
	http.ListenAndServe(":8080", nil)
}
