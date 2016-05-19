package main

import (
	"net/http"
	log "github.com/cihub/seelog"
)

func LiveHandler(w http.ResponseWriter, r *http.Request){
	log.Debug(r.Method)
	buf := make([]byte, 10240)
	for{
		len, err := r.Body.Read(buf)
		if err !=nil{
			log.Debug(len)
			log.Error(err)
			break;
		}
		log.Debug(len)
	}

}

func main(){
	http.HandleFunc("/live/" , LiveHandler)

	http.ListenAndServe(":8888", nil)
}
