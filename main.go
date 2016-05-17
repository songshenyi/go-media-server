package main

import "net/http"

func LiveHandler(w http.ResponseWriter, r *http.Request){

}

func main(){
	http.HandleFunc("/live/" , LiveHandler)

	http.ListenAndServe(":8888", nil)
}
