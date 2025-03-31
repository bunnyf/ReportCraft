# ReportCraft Go测试床

这个测试床用于批量运行和测试ReportCraft的各种示例配置，验证其功能正确性并测量性能。

## 功能特点

- 批量测试多个配置文件
- 支持并行测试以提高效率
- 记录每个测试的执行时间和结果
- 支持多种输出格式（表格、CSV、JSON）
- 可控制的并发度

## 使用方法

### 基本用法

```bash
# 在默认位置测试所有JSON配置文件
go run main.go

# 指定ReportCraft可执行文件路径和示例目录
go run main.go -path=/path/to/reportcraft -examples=/path/to/examples
```

### 高级选项

```bash
# 只测试特定模式的配置文件
go run main.go -patterns="spectrum*.json,waveform*.json"

# 并行执行4个测试
go run main.go -parallel=4

# 输出CSV格式并保存到文件
go run main.go -format=csv -output=results.csv
```

### 所有选项

| 选项 | 默认值 | 说明 |
|------|--------|------|
| `-path` | `../../reportcraft` | ReportCraft可执行文件路径 |
| `-examples` | `../../examples` | 示例配置文件目录 |
| `-patterns` | `*.json` | 要测试的配置文件模式，多个模式用逗号分隔 |
| `-parallel` | `1` | 并行执行的测试数量 |
| `-format` | `table` | 输出格式: table, csv, json |
| `-output` | 空（输出到控制台） | 测试结果输出文件路径 |

## 示例输出

### 表格输出

```
配置文件                       | 执行时间         | 输出路径                                            | 状态    
------------------------------------------------------------------------------------------------------
spectrum-report-embedded.json | 1.235s          | ./output/spectrum-analysis-embedded.html            | 成功    
waveform-multi-channel.json   | 875.421ms       | ./output/waveform-multi-channel.html               | 成功    
trend-analysis-multi.json     | 452.125ms       | ./output/trend-analysis-multi.docx                 | 成功    
```

### JSON输出

```json
[
  {
    "configFile": "spectrum-report-embedded.json",
    "execTimeMs": 1235,
    "outputPath": "./output/spectrum-analysis-embedded.html",
    "status": "成功"
  },
  {
    "configFile": "waveform-multi-channel.json",
    "execTimeMs": 875,
    "outputPath": "./output/waveform-multi-channel.html",
    "status": "成功"
  }
]
```

## 构建可执行文件

```bash
go build -o rc-testbed
```

## 使用场景

1. 功能测试：验证所有示例配置是否能正确生成报告
2. 性能测试：测量不同类型报告的生成时间
3. 回归测试：确保代码变更不会破坏现有功能
4. 批量生成：一次性生成所有报告用于展示或演示
