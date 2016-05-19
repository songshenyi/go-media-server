package main

import (
	"net/http"
	log "github.com/cihub/seelog"
)

func LiveHandler(w http.ResponseWriter, r *http.Request){
	log.Debug(r.Method)
}

func main(){
	http.HandleFunc("/live/" , LiveHandler)

	http.ListenAndServe(":8888", nil)
}
