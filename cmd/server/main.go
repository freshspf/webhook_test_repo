package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"webhook-demo/internal/handlers"
)

func main() {
	r := mux.NewRouter()
	
	// GitHub webhook 路由
	r.HandleFunc("/webhook", handlers.WebhookHandler).Methods("POST")
	
	// 俄罗斯方块游戏路由
	r.HandleFunc("/tetris", handlers.TetrisHandler).Methods("GET")
	r.HandleFunc("/", handlers.TetrisHandler).Methods("GET") // 默认页面也是俄罗斯方块
	
	// 五子棋游戏路由
	r.HandleFunc("/gomoku", handlers.GomokuHandler).Methods("GET")
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Server starting on port %s", port)
	log.Printf("Tetris game available at: http://localhost:%s/tetris", port)
	log.Printf("Gomoku game available at: http://localhost:%s/gomoku", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}