package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yuxxeun/VG/internal/router"
)

func main() {
	port := "8080"

	http.HandleFunc("/", router.Route)

	fmt.Printf("🚀 Server running on http://localhost:%s\n\n", port)
	fmt.Println("📍 Available endpoints:")
	fmt.Println("   GET  http://localhost:" + port + "/")
	fmt.Println("   GET  http://localhost:" + port + "/health")
	fmt.Println("\n📍 API V1:")
	fmt.Println("   GET  http://localhost:" + port + "/api/v1/users")
	fmt.Println("   POST http://localhost:" + port + "/api/v1/users")
	fmt.Println("   GET  http://localhost:" + port + "/api/v1/users/:id")
	fmt.Println("   GET  http://localhost:" + port + "/api/v1/products")
	fmt.Println("   POST http://localhost:" + port + "/api/v1/products")
	fmt.Println("\n📍 API V2:")
	fmt.Println("   GET  http://localhost:" + port + "/api/v2/users")
	fmt.Println()

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
