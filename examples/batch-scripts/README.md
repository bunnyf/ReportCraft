# ReportCraft 批处理脚本

这些批处理脚本用于自动批量生成ReportCraft报告，支持Windows和Linux/macOS平台。

## 脚本功能

- 自动查找和处理所有JSON配置文件
- 批量生成多种格式的报告
- 提供成功/失败的详细日志
- 自动识别输出格式
- 支持自定义输出目录

## 使用方法

### Windows

```batch
# 进入批处理脚本目录
cd examples/batch-scripts

# 运行批处理脚本
generate-reports.bat
```

### Linux/macOS

```bash
# 进入批处理脚本目录
cd examples/batch-scripts

# 添加执行权限
chmod +x generate-reports.sh

# 运行脚本
./generate-reports.sh
```

## 脚本参数定制

可以根据需要修改脚本中的以下参数：

| 参数 | 说明 |
|------|------|
| `REPORTCRAFT_PATH` | ReportCraft可执行文件路径 |
| `EXAMPLES_DIR` | 配置文件目录 |
| `OUTPUT_DIR` | 输出报告的目录 |
| `LOG_FILE` | 日志文件路径 |

## 输出

脚本执行后会产生以下输出：

1. 控制台输出：显示每个报告的处理进度和结果
2. 日志文件：详细记录所有报告生成的信息和错误
3. 生成的报告：保存在指定的输出目录中

## 日志格式

日志文件格式如下：

```text
ReportCraft批量生成日志
生成时间: 2025-03-20 10:15:30

[成功] spectrum-3d
  配置: ../spectrum-3d.json
  输出: ../output/batch/spectrum-3d.html
  报告生成成功...

[失败] invalid-report
  配置: ../invalid-report.json
  错误:
  错误: 无法解析配置文件...

摘要:
总共: 10, 成功: 9, 失败: 1
```

## 故障排除

1. **ReportCraft可执行文件未找到**
   - 检查可执行文件路径是否正确
   - 确保已构建ReportCraft

2. **权限错误**
   - 对于Linux/macOS用户，确保脚本和ReportCraft具有执行权限

3. **配置错误**
   - 检查日志文件中的详细错误信息
   - 验证JSON配置文件格式是否正确

## 集成到自动化流程

这些脚本可以集成到CI/CD流程中，实现自动报告生成：

```yaml
# 示例CI配置片段
build:
  script:
    - cd examples/batch-scripts
    - chmod +x generate-reports.sh
    - ./generate-reports.sh
  artifacts:
    paths:
      - examples/output/**/*
