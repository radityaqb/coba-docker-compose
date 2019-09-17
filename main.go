package main

import "net/http"

func main() {
	initRedis()
	initDB()

	http.HandleFunc("/post_data", handlePostData)
	http.HandleFunc("/get_data", handleGetData)

	http.ListenAndServe(":8181", nil)
}

func handlePostData(w http.ResponseWriter, r *http.Request) {
	// Create func to insert DB and insert Redis
	w.Write([]byte("success"))
}

func handleGetData(w http.ResponseWriter, r *http.Request) {
	// Create func to get Redis by ID
	// if redis is empty, get from db
	w.Write([]byte("success"))
}
