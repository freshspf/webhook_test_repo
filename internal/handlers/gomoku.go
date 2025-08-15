package handlers

import (
	"net/http"
	"os"
	"path/filepath"
)

func GomokuHandler(w http.ResponseWriter, r *http.Request) {
	htmlFile := filepath.Join("web", "gomoku.html")
	
	content, err := os.ReadFile(htmlFile)
	if err != nil {
		http.Error(w, "五子棋游戏页面加载失败", http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(content)
}