package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/auth", authorize)
	http.ListenAndServe(":9099", nil)
}

const auth = "Authorization"
const originAuth = "jialinwu"

func authorize(w http.ResponseWriter, r *http.Request) {
	log.Printf("r.Head: %+v\n", r.Header)
	key := r.Header.Get(auth)
	if len(key) < 1 {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	mCoder := md5.New()
	mCoder.Write([]byte(originAuth))

	// o := fmt.Sprintf("%x", mCoder.Sum(nil))
	// log.Printf("o: %+v\n", o)
	o := hex.EncodeToString(mCoder.Sum(nil)) //hex.EncodeToString is faster one time than fmt.Sprintf("%x")
	log.Printf("2o: %+v\n", o)
	// o := base64.NewEncoding(originAuth)
	log.Printf("key: %+v\n", key)
	// w.Write([]byte("23"))
	w.WriteHeader(http.StatusOK)
	// w.Write([]byte("23"))
}
