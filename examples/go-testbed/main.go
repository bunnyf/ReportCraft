package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type TestResult struct {
	ConfigName  string
	ExecTime    time.Duration
	OutputPath  string
	Success     bool
	ErrorOutput string
}

func main() {
	// 定义命令行参数
	rcPath := flag.String("path", "../../reportcraft", "ReportCraft可执行文件路径")
	examplesDir := flag.String("examples", "../../examples", "示例目录")
	patterns := flag.String("patterns", "*.json", "要测试的配置文件模式，多个模式用逗号分隔")
	parallel := flag.Int("parallel", 1, "并行执行的测试数量")
	outputFormat := flag.String("format", "table", "输出格式: table, csv, json")
	outputPath := flag.String("output", "", "测试结果输出文件路径（如果未指定则输出到控制台）")
	flag.Parse()

	// 检查可执行文件是否存在
	if _, err := os.Stat(*rcPath); os.IsNotExist(err) {
		log.Fatalf("错误: ReportCraft可执行文件不存在于路径 %s", *rcPath)
	}

	// 获取所有匹配的配置文件
	var configFiles []string
	patternList := strings.Split(*patterns, ",")
	for _, pattern := range patternList {
		matches, err := filepath.Glob(filepath.Join(*examplesDir, pattern))
		if err != nil {
			log.Fatalf("错误: 无法查找示例文件: %v", err)
		}
		configFiles = append(configFiles, matches...)
	}

	if len(configFiles) == 0 {
		log.Fatalf("错误: 没有找到匹配的配置文件")
	}

	fmt.Printf("将测试 %d 个配置文件，并行数量: %d\n", len(configFiles), *parallel)
	fmt.Println("开始测试...")

	// 运行测试
	results := runTests(configFiles, *rcPath, *parallel)

	// 输出结果
	outputResults(results, *outputFormat, *outputPath)

	// 输出摘要
	successCount := 0
	for _, r := range results {
		if r.Success {
			successCount++
		}
	}
	fmt.Printf("\n测试摘要:\n")
	fmt.Printf("总共: %d, 成功: %d, 失败: %d\n", len(results), successCount, len(results)-successCount)
}

func runTests(configFiles []string, rcPath string, parallel int) []TestResult {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	results := make([]TestResult, 0, len(configFiles))

	// 创建信号量控制并发
	semaphore := make(chan struct{}, parallel)

	for _, configFile := range configFiles {
		wg.Add(1)
		semaphore <- struct{}{} // 获取信号量

		go func(cf string) {
			defer wg.Done()
			defer func() { <-semaphore }() // 释放信号量

			// 执行命令并计时
			configName := filepath.Base(cf)
			fmt.Printf("测试 %s...\n", configName)

			start := time.Now()
			cmd := exec.Command(rcPath, "-config="+cf)
			output, err := cmd.CombinedOutput()
			execTime := time.Since(start)

			// 解析输出以获取生成的报告路径
			outputPath := "未知"
			outputLines := strings.Split(string(output), "\n")
			for _, line := range outputLines {
				if strings.Contains(line, "Generated") && strings.Contains(line, "at") {
					parts := strings.Split(line, "at")
					if len(parts) > 1 {
						outputPath = strings.TrimSpace(parts[1])
						break
					}
				}
			}

			// 创建测试结果
			result := TestResult{
				ConfigName:  configName,
				ExecTime:    execTime,
				OutputPath:  outputPath,
				Success:     err == nil,
				ErrorOutput: string(output),
			}

			// 线程安全地添加结果
			mutex.Lock()
			results = append(results, result)
			mutex.Unlock()
		}(configFile)
	}

	wg.Wait()
	return results
}

func outputResults(results []TestResult, format string, outputPath string) {
	var output string

	switch format {
	case "table":
		output = formatAsTable(results)
	case "csv":
		output = formatAsCSV(results)
	case "json":
		output = formatAsJSON(results)
	default:
		output = formatAsTable(results)
	}

	// 输出结果
	if outputPath == "" {
		fmt.Println(output)
	} else {
		err := os.WriteFile(outputPath, []byte(output), 0644)
		if err != nil {
			log.Fatalf("错误: 无法写入输出文件: %v", err)
		}
		fmt.Printf("结果已保存到: %s\n", outputPath)
	}
}

func formatAsTable(results []TestResult) string {
	var sb strings.Builder

	// 表头
	sb.WriteString(fmt.Sprintf("%-30s | %-15s | %-50s | %-7s\n", "配置文件", "执行时间", "输出路径", "状态"))
	sb.WriteString(fmt.Sprintf("%s\n", strings.Repeat("-", 110)))

	// 表内容
	for _, r := range results {
		status := "成功"
		if !r.Success {
			status = "失败"
		}
		sb.WriteString(fmt.Sprintf("%-30s | %-15s | %-50s | %-7s\n",
			r.ConfigName,
			r.ExecTime.Round(time.Millisecond).String(),
			r.OutputPath,
			status))
	}

	return sb.String()
}

func formatAsCSV(results []TestResult) string {
	var sb strings.Builder

	// CSV 头
	sb.WriteString("配置文件,执行时间(ms),输出路径,状态\n")

	// CSV 内容
	for _, r := range results {
		status := "成功"
		if !r.Success {
			status = "失败"
		}
		sb.WriteString(fmt.Sprintf("%s,%d,%s,%s\n",
			r.ConfigName,
			r.ExecTime.Milliseconds(),
			r.OutputPath,
			status))
	}

	return sb.String()
}

func formatAsJSON(results []TestResult) string {
	var sb strings.Builder
	sb.WriteString("[\n")

	for i, r := range results {
		status := "成功"
		if !r.Success {
			status = "失败"
		}

		sb.WriteString(fmt.Sprintf(`  {
    "configFile": "%s",
    "execTimeMs": %d,
    "outputPath": "%s",
    "status": "%s"
  }`, r.ConfigName, r.ExecTime.Milliseconds(), r.OutputPath, status))

		if i < len(results)-1 {
			sb.WriteString(",")
		}
		sb.WriteString("\n")
	}

	sb.WriteString("]\n")
	return sb.String()
}
