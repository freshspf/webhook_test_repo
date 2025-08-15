package handlers

import (
	"net/http"
	"path/filepath"
)

// TetrisHandler 处理俄罗斯方块游戏页面请求
func TetrisHandler(w http.ResponseWriter, r *http.Request) {
	// 设置内容类型为HTML
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	
	// 获取HTML文件路径
	htmlPath := filepath.Join("web", "tetris.html")
	
	// 提供静态文件服务
	http.ServeFile(w, r, htmlPath)
}