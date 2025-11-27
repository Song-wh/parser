package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const PORT = 9990

func main() {
	// 핸들러 설정
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// CORS 헤더 추가
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")

		// 루트 경로는 index.html 서빙
		urlPath := r.URL.Path
		if urlPath == "/" {
			urlPath = "/index.html"
		}

		// 파일 경로 생성
		filePath := filepath.Join(".", urlPath)

		// 파일 존재 확인
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			http.NotFound(w, r)
			log.Printf("404 %s %s", r.Method, urlPath)
			return
		}

		log.Printf("%s %s", r.Method, urlPath)
		http.ServeFile(w, r, filePath)
	})

	fmt.Println("\n========================================")
	fmt.Printf("🚀 서버가 시작되었습니다!\n")
	fmt.Printf("📡 http://localhost:%d\n", PORT)
	fmt.Println("========================================")
	fmt.Println("종료하려면 Ctrl+C를 누르세요.\n")

	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil); err != nil {
		log.Fatal(err)
	}
}
