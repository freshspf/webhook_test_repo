package services

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/webhook-demo/internal/models"
)

// EventProcessor 事件处理器
type EventProcessor struct {
	githubService     *GitHubService
	claudeCodeService *ClaudeCodeCLIService
	gitService        *GitService
	commandRegex      *regexp.Regexp
}

// NewEventProcessor 创建新的事件处理器
func NewEventProcessor(githubService *GitHubService, claudeCodeService *ClaudeCodeCLIService, gitService *GitService) *EventProcessor {
	return &EventProcessor{
		githubService:     githubService,
		claudeCodeService: claudeCodeService,
		gitService:        gitService,
		commandRegex:      regexp.MustCompile(`^/(code|continue|fix|help|report)\s*(.*)$`),
	}
}

// executeCommand 执行命令
func (ep *EventProcessor) executeCommand(command *Command, ctx *CommandContext) error {
	log.Printf("执行命令: %s, 参数: %s", command.Command, command.Args)

	switch command.Command {
	case "code":
		return ep.handleCodeCommand(command, ctx)
	case "continue":
		return ep.handleContinueCommand(command, ctx)
	case "fix":
		return ep.handleFixCommand(command, ctx)
	case "help":
		return ep.handleHelpCommand(command, ctx)
	case "report":
		return ep.handleReportCommand(command, ctx)
	default:
		return fmt.Errorf("未知命令: %s", command.Command)
	}
}

// handleReportCommand 处理报告生成命令
func (ep *EventProcessor) handleReportCommand(command *Command, ctx *CommandContext) error {
	log.Printf("生成仓库活动报告")

	// 分解命令参数，支持指定天数
	days := 7 // 默认7天
	if command.Args != "" {
		fmt.Sscanf(command.Args, "%d", &days)
		if days <= 0 {
			days = 7
	}
	}

	repo := strings.Split(ctx.Repository.FullName, "/")
	if len(repo) != 2 {
		return fmt.Errorf("无效的仓库名称: %s", ctx.Repository.FullName)
	}
	owner, repoName := repo[0], repo[1]

	// 获取最近的活动数据
	issues, _ := ep.githubService.GetRecentIssues(owner, repoName, days)
	prs, _ := ep.githubService.GetRecentPullRequests(owner, repoName, days)
	commits, _ := ep.githubService.GetRecentCommits(owner, repoName, days)

	// 生成报告
	report := fmt.Sprintf(`📊 **%d天活动报告**

### 📝 Issues (%d)
`, days, len(issues))

	// 添加Issue统计
	for _, issue := range issues {
		report += fmt.Sprintf("- #%d: %s (%s)\n", issue.Number, issue.Title, issue.State)
	}

	report += fmt.Sprintf("\n### 🔀 Pull Requests (%d)\n", len(prs))

	// 添加PR统计
	for _, pr := range prs {
		report += fmt.Sprintf("- #%d: %s (%s)\n", pr.Number, pr.Title, pr.State)
	}

	report += fmt.Sprintf("\n### 💾 最近提交 (%d)\n", len(commits))

	// 添加提交记录
	for _, commit := range commits {
		report += fmt.Sprintf("- %.7s: %s\n", commit.SHA, commit.Message)
	}

	report += fmt.Sprintf("\n---\n*生成时间: %s*", time.Now().Format("2006-01-02 15:04:05"))

	return ep.createResponse(ctx, report)
}

// handleHelpCommand 处理帮助命令
func (ep *EventProcessor) handleHelpCommand(command *Command, ctx *CommandContext) error {
	log.Printf("处理帮助命令")

	response := `📖 **CodeAgent 帮助**

**支持的命令:**

🔹 ` + "`" + `/code <需求描述>` + "`" + ` - 自动分析并实现到代码库
🔹 ` + "`" + `/continue [说明]` + "`" + ` - 继续当前的开发任务
🔹 ` + "`" + `/fix <问题描述>` + "`" + ` - 修复指定的代码问题
🔹 ` + "`" + `/report [天数]` + "`" + ` - 生成仓库活动报告
🔹 ` + "`" + `/help` + "`" + ` - 显示此帮助信息

**使用示例:**
- ` + "`" + `/code 创建一个用户登录API` + "`" + ` - 自动分析并实现到项目中
- ` + "`" + `/code 添加JWT认证功能` + "`" + ` - 自动分析并修改代码
- ` + "`" + `/continue 添加数据验证逻辑` + "`" + `
- ` + "`" + `/fix 修复空指针异常` + "`" + `
- ` + "`" + `/report 30` + "`" + ` - 生成最近30天的活动报告

**工作流程:**
1. 🎯 在Issue或PR评论中输入命令
2. 🤖 AI分析需求并生成代码
3. 🌲 创建独立的Git工作空间
4. 📝 自动提交代码并创建PR
5. 💬 在GitHub界面展示结果

---
*GitHub Webhook Demo v1.0*`

	return ep.createResponse(ctx, response)
}

[... rest of the file remains unchanged ...]