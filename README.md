# GitHub Webhook Demo with Tetris Game

这是一个集成了俄罗斯方块游戏的GitHub Webhook演示项目。

## 功能特性

### GitHub Webhook 处理
- 自动处理GitHub的各种事件
- 支持Issue事件的自动化处理
- 集成Claude Code CLI进行智能代码修改
- 自动创建分支和Pull Request

### 俄罗斯方块游戏
- 经典的俄罗斯方块游戏体验
- 完整的游戏逻辑和计分系统
- 响应式设计，支持键盘控制
- 实时显示分数、等级和消除行数

## 快速开始

### 安装依赖
```bash
go mod download
```

### 运行服务器
```bash
go run cmd/server/main.go
```

服务器启动后，访问以下URL：
- 俄罗斯方块游戏: `http://localhost:8080/tetris`
- GitHub Webhook端点: `http://localhost:8080/webhook`

## 游戏控制说明

- **← →** 左右移动方块
- **↓** 快速下降
- **↑** 旋转方块
- **空格** 暂停/继续游戏

## 项目结构

```
.
├── cmd/server/main.go          # 主服务器入口
├── internal/
│   ├── handlers/
│   │   ├── webhook.go          # GitHub Webhook处理器
│   │   └── tetris.go           # 俄罗斯方块游戏处理器
│   ├── models/                 # 数据模型
│   └── services/               # 业务逻辑服务
├── web/
│   └── tetris.html             # 俄罗斯方块游戏页面
└── README.md
```

## 环境变量

- `PORT`: 服务器端口（默认: 8080）
- `GITHUB_TOKEN`: GitHub访问令牌
- `GITHUB_WEBHOOK_SECRET`: Webhook密钥

## 部署说明

### 本地开发
1. 克隆仓库
2. 设置必要的环境变量
3. 运行 `go run cmd/server/main.go`

### 生产部署
1. 构建二进制文件: `go build -o webhook-server cmd/server/main.go`
2. 设置生产环境变量
3. 运行服务器

## 贡献指南

欢迎提交Issue和Pull Request来改进这个项目。

## 许可证

MIT License