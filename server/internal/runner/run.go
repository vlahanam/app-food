package runner

import (
	"fmt"
	"net/http"
	"os"
)

// [VN] Run khởi chạy ứng dụng web server
func Run() {
	fmt.Println("Long An - Starting App Food Server...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// [VN] Endpoint kiểm tra sức khỏe ứng dụng
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok","service":"app-food"}`))
	})

	// [VN] Endpoint API kiểm tra sức khỏe
	http.HandleFunc("/api/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok","service":"app-food-api"}`))
	})

	// [VN] Endpoint mặc định
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message":"Welcome to App Food API (Golang 1.26)"}`))
	})

	fmt.Printf("App Food Server is listening on port %s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Server failed: %v\n", err)
	}
}
