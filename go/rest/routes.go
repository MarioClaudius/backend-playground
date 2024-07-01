package rest

// import (
// 	"fmt"
// 	"net/http"
// )

// func GetHandler() *http.ServeMux {
// 	handler := http.NewServeMux()
// 	initRoute(handler)

// 	return handler
// }

// func initRoute(h *http.ServeMux) {
// 	// example of anonymous function
// 	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello World")
// 	})
// 	h.HandleFunc("/detail", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, r.Method)
// 		fmt.Fprintln(w, r.URL)
// 	})
// }
