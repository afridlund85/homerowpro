package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, time.Now().Format("02 Jan 2006 15:04:05 MST"))
}

func main() {
	pub_dir := "web/public"
	_, err := os.Stat(pub_dir)
	if os.IsNotExist(err) {
		fmt.Printf("Dir '%s' does not exist.", pub_dir)
		return
	}
	httpHandler := http.FileServer(http.Dir(pub_dir))
	http.HandleFunc("/time", timeHandler)
	http.Handle("/", httpHandler)

	fmt.Println("Server started on 8080")
	http.ListenAndServe(":8080", nil)
}
