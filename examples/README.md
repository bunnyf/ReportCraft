# ReportCraft 功能示例清单

本目录包含各种ReportCraft功能的示例和演示，帮助用户快速了解和使用ReportCraft的强大功能。

## 功能分类

ReportCraft支持以下主要功能：

1. **基础报告**
   - Hello World文本报告 (`hello-word.json`)
   - Hello Excel表格报告 (`hello-excel.json`)
   - 使用模板的报告 (`hello-template.json`)

2. **表格报告**
   - 使用嵌入数据的表格报告 (`table-report-embedded.json`)
   - 使用CSV数据的表格报告 (`table-report-csv.json`)
   - 使用JSON数据的表格报告 (`table-report-json.json`)

3. **趋势分析报告**
   - 基本趋势分析 (`trend-report.json`)
   - 多时间序列趋势 (`trend-analysis-multi.json`)
   - 预测趋势 (`trend-forecast.json`)

4. **波形分析报告**
   - 基本波形分析 (`waveform-report.json`)
   - 多通道波形 (`waveform-multi-channel.json`)
   - 波形对比分析 (`waveform-comparison.json`)

5. **频谱分析报告**
   - 使用嵌入数据的频谱分析 (`spectrum-report-embedded.json`)
   - 使用CSV数据的频谱分析 (`spectrum-report-csv.json`)
   - 高级3D频谱 (`spectrum-3d.json`)
   - 瀑布图分析 (`spectrum-waterfall.json`)

6. **高级图表报告**
   - 组合图表报告 (`combined-charts.json`)
   - 交互式仪表盘 (`interactive-dashboard.json`)
   - 数据过滤与分组 (`data-filtering.json`)

7. **演示工具**
   - Go测试床 (`go-testbed/`)
   - Python接口 (`python-interface/`)
   - Windows应用接口 (`windows-app/`)

## 目录结构

- `*.json` - 各种类型的报告配置文件示例
- `data/` - 示例数据文件
- `batch-scripts/` - 批量生成报告的脚本
- `go-testbed/` - Go语言测试床
- `python-interface/` - Python接口

## 详细示例列表

### 基础报告

- **Hello World 文本报告** - 最简单的文本报告示例
- **Hello Excel 表格报告** - 基础Excel表格报告
- **基于模板的报告** - 使用预定义HTML/Docx模板的报告

### 表格报告

- **嵌入数据表格报告** - 直接在配置中嵌入数据的表格报告
- **CSV数据表格报告** - 从CSV文件导入数据的表格报告
- **JSON数据表格报告** - 从JSON文件导入数据的表格报告

### 趋势分析报告

- **多参数趋势分析** - [trend-analysis-multi.json](trend-analysis-multi.json) - 展示多个参数随时间变化的趋势
- **预测趋势分析** - [trend-forecast.json](trend-forecast.json) - 包含预测功能的趋势分析

### 波形分析报告

- **多通道波形分析** - [waveform-multi-channel.json](waveform-multi-channel.json) - 多个采集通道的波形分析
- **波形对比分析** - [waveform-comparison.json](waveform-comparison.json) - 不同条件下的波形对比分析

### 频谱分析报告

- **3D频谱分析** - [spectrum-3d.json](spectrum-3d.json) - 三维表面图的频谱分析
- **瀑布图谱分析** - [spectrum-waterfall.json](spectrum-waterfall.json) - 瀑布图形式的频谱分析
- **增强瀑布图谱分析** - [spectrum-waterfall-enhanced.json](spectrum-waterfall-enhanced.json) - 添加高级视觉效果的瀑布图分析

### 高级图表报告

- **组合图表报告** - [combined-charts.json](combined-charts.json) - 在同一报告中组合多种图表类型
- **综合图表仪表板** - [combined-charts-comprehensive.json](combined-charts-comprehensive.json) - 全面的多图表分析仪表板
- **交互式仪表板** - [interactive-dashboard.json](interactive-dashboard.json) - 提供交互功能的实时监控仪表板
- **高级交互式仪表板** - [interactive-dashboard-advanced.json](interactive-dashboard-advanced.json) - 带有高级交互控件和多视图的监控仪表板

## 工具与接口

### Go测试床

[go-testbed](go-testbed/) 目录包含用于批量测试ReportCraft功能的Go应用程序。支持并行测试、自定义输出格式和详细日志记录。

#### Go测试床功能

- 批量测试多个配置文件
- 支持并行测试以提高效率
- 记录每个测试的执行时间和结果
- 支持多种输出格式（表格、CSV、JSON）

### Python接口

[python-interface](python-interface/) 目录提供了从Python程序调用ReportCraft的接口。便于集成到现有的Python应用程序或数据分析流程中。

#### Python接口功能

- 简单易用的API，隐藏底层复杂性
- 支持所有ReportCraft报表类型和配置
- 提供便捷的报表类型专用方法
- 支持命令行和程序化使用

### 批处理脚本

[batch-scripts](batch-scripts/) 目录包含用于批量生成报告的脚本，支持Windows和Linux/macOS平台。

#### 批处理脚本功能

- 自动查找和处理所有JSON配置文件
- 批量生成多种格式的报告
- 提供成功/失败的详细日志
- 支持自定义输出目录

## 示例数据

所有示例数据位于 [data](data/) 目录中：

- **多参数趋势数据** - [multi_trend_data.csv](data/multi_trend_data.csv)
- **预测趋势数据** - [forecast_data.csv](data/forecast_data.csv)
- **多通道波形数据** - [multichannel_waveform.csv](data/multichannel_waveform.csv)
- **波形对比数据** - [comparison_waveform.csv](data/comparison_waveform.csv)
- **3D频谱数据** - [3d_spectrum_data.csv](data/3d_spectrum_data.csv)
- **瀑布图谱数据** - [waterfall_data.csv](data/waterfall_data.csv)
- **组合图表数据** - [combined_charts_data.json](data/combined_charts_data.json)
- **仪表板数据** - [dashboard_data.json](data/dashboard_data.json)

## 运行示例

每个示例都可以通过以下方式运行：

```bash
# 使用命令行
reportcraft -config=examples/spectrum-3d.json

# 使用批处理脚本
cd examples/batch-scripts
./generate-reports.sh   # Linux/macOS
generate-reports.bat    # Windows

# 使用Go测试床
cd examples/go-testbed
go run main.go

# 使用Python接口
cd examples/python-interface
python examples.py
```

## 最佳实践

### 3D频谱分析优化

为获得最佳的3D频谱分析效果，建议使用以下配置：

```json
"chartOptions": {
  "colorOpacity": 0.85,
  "meshDensity": {"x": 50, "y": 50},
  "lighting": {
    "ambient": 0.65,
    "diffuse": 0.85,
    "specular": 0.15
  },
  "interactiveControls": true
}
```

### 瀑布图优化

为获得最佳的瀑布图显示效果，建议使用以下配置：

```json
"chartOptions": {
  "gradientScheme": "rainbow",
  "lineWidth": 1.2,
  "fillOpacity": 0.7,
  "showGridLines": true,
  "gridColor": "#dddddd",
  "yAxisMode": "time",
  "renderQuality": "high"
}
```

## 常见问题

**Q: 如何自定义报告输出路径？**

A: 在配置文件中设置 `outputPath` 字段，或使用命令行参数 `-output=路径`

**Q: 如何在一个报告中展示多种图表？**

A: 使用 `combined-report` 类型，并在 `sections` 中定义多个图表

**Q: 如何提高3D图表的渲染质量？**

A: 设置 `meshDensity` 属性和 `renderQuality: "high"`

**Q: 如何批量生成多个报告？**

A: 使用 `batch-scripts` 目录中的批处理脚本或 `go-testbed` 测试工具
