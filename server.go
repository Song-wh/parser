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
	// 현재 디렉토리를 정적 파일 서빙 루트로 설정
	fs := http.FileServer(http.Dir("."))
	
	// 핸들러 래핑 (로깅 및 CORS 추가)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 루트 경로는 index.html로 리다이렉트
		if r.URL.Path == "/" {
			r.URL.Path = "/index.html"
		}
		
		// CORS 헤더 추가
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
		
		// 파일 존재 확인
		path := filepath.Join(".", r.URL.Path)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			http.NotFound(w, r)
			log.Printf("404 %s %s", r.Method, r.URL.Path)
			return
		}
		
		log.Printf("%s %s", r.Method, r.URL.Path)
		fs.ServeHTTP(w, r)
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

