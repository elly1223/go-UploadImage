package main

import (
	"fmt"
	"net/http"
	"os"
)

func GetImage(w http.ResponseWriter, r *http.Request) {
	const testPath = "../image00001.jpg"
	data, err := os.ReadFile(testPath)
	// READ - FAILURE
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// READ - SUCCESS
	fmt.Println(r.RemoteAddr, len(data))
	w.WriteHeader(http.StatusAccepted)
	w.Header().Add("content-type", "image/jpeg")
	w.Write(data)
}

func main() {
	http.HandleFunc("/api/3", GetImage)
	fmt.Printf("listening... port: 80\n")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Printf("ERR: http.ListenAndServe: %s\n", err)
		return
	}
}
