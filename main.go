package main

import (
	"fmt"
	"log"
	"net/http"

	handler "github.com/yuxxeun/VG/api/api"
)

func main() {
	port := "8080"

	http.HandleFunc("/", handler.Handler)

	fmt.Printf("Server running on http://localhost:%s\n", port)
	fmt.Println("Available endpoints:")
	fmt.Println("  GET  http://localhost:" + port + "/")
	fmt.Println("  GET  http://localhost:" + port + "/api/users")
	fmt.Println("  POST http://localhost:" + port + "/api/users")
	fmt.Println("  GET  http://localhost:" + port + "/api/health")

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
