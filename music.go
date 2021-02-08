package main

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func (app *App) HandleMusicUpload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		file, handler, err := r.FormFile("file")
		if err != nil {
			panic(err)
		}
		defer file.Close()
		rn := rand.New(rand.NewSource(99))
		code := rn.Uint64()
		stringified_code := strconv.FormatUint(code, 10)
		address := "/home/vader/phoenix-music/" + stringified_code + ".mp3"
		log.Println(address)
		f, err := os.Create(address)
		if err != nil {
			log.Println(err)
			res, _ := json.Marshal(map[string]interface{}{
				"error": "cannot upload the file.",
			})
			w.Write(res)
			return
		}
		defer f.Close()
		_, _ = io.Copy(f, file)
		app.db.Create(&Music{
			Name:     r.PostForm.Get("name"),
			code:     code,
			FileName: handler.Filename,
		})
		res, _ := json.Marshal(map[string]interface{}{
			"status": "ok",
			"code":   code,
		})
		w.Write(res)
	}
}

func (app *App) HandleStreamMusic() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Connection", "Keep-Alive")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Transfer-Encoding", "chunked")
		w.Header().Set("Content-Type", "audio/mpeg")
		filename := r.URL.Query()["music"][0]
		address := "/home/vader/phoenix-music/" + filename + ".mp3"
		f, err := os.Open(address)
		if err != nil {
			res, _ := json.Marshal(map[string]interface{}{
				"error": "no such music found",
			})
			w.Write(res)
			return
		}
		for {
			b := make([]byte, 44100)
			_, err = f.Read(b)
			if err == io.EOF {
				break
			}
			w.Write(b)
		}
	}
}
