package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/webhook-demo/internal/config"
	"github.com/webhook-demo/internal/handlers"
	"github.com/webhook-demo/internal/middleware"
	"github.com/webhook-demo/internal/services"
)

// isNarcissistic checks if a number is a narcissistic number
func isNarcissistic(num int) bool {
	if num <= 0 {
		return false
	}

	// Convert number to string to count digits
	numStr := fmt.Sprintf("%d", num)
	power := len(numStr)
	
	// Calculate sum of digits raised to power
	sum := 0
	temp := num
	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(power)))
		temp /= 10
	}
	
	return sum == num
}

// findNarcissisticNumbers finds all narcissistic numbers in the given range
func findNarcissisticNumbers(start, end int) []int {
	if start > end || start < 0 {
		return nil
	}

	var result []int
	for i := start; i <= end; i++ {
		if isNarcissistic(i) {
			result = append(result, i)
		}
	}
	return result
}

func main() {
	// Find narcissistic numbers from 1 to 500
	narcissisticNumbers := findNarcissisticNumbers(1, 500)
	fmt.Println("Narcissistic numbers between 1 and 500:")
	fmt.Println(narcissisticNumbers)

	// 加载配置
	cfg := config.Load()

	// 初始化服务
	githubService := services.NewGitHubService(cfg.GitHub.Token)
	claudeCodeService := services.NewClaudeCodeCLIService(&cfg.ClaudeCodeCLI)
	gitConfig := config.LoadGitConfig()
	gitService := services.NewGitService(gitConfig.WorkDir)
	eventProcessor := services.NewEventProcessor(githubService, claudeCodeService, gitService)

	// 初始化处理器
	webhookHandler := handlers.NewWebhookHandler(eventProcessor, cfg.GitHub.WebhookSecret)

	// 设置路由
	router := setupRouter(webhookHandler, cfg)

	// 启动服务器
	srv := &http.Server{
		// Addr:    ":" + cfg.Server.Port,
		Addr:    "0.0.0.0:" + cfg.Server.Port,
		Handler: router,
	}

	// 优雅关闭
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("启动服务器失败: %v", err)
		}
	}()

	log.Printf("Webhook服务器已启动，端口: %s", cfg.Server.Port)
	log.Printf("Webhook端点: http://localhost:%s/webhook", cfg.Server.Port)

	// 等待中断信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("正在关闭服务器...")

	// 优雅关闭超时
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("服务器强制关闭:", err)
	}

	log.Println("服务器已退出")
}

func setupRouter(webhookHandler *handlers.WebhookHandler, cfg *config.Config) *gin.Engine {
	if cfg.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	// 中间件
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORS())

	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"timestamp": time.Now().Unix(),
		})
	})

	// Webhook端点
	router.POST("/webhook", webhookHandler.HandleWebhook)

	// API信息
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name":        "GitHub Webhook Demo",
			"version":     "1.0.0",
			"description": "GitHub Webhook处理演示",
			"endpoints": map[string]string{
				"webhook": "/webhook",
				"health":  "/health",
			},
		})
	})

	return router
}