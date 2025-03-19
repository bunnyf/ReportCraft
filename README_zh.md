# GenRep - 通用报表生成工具

GenRep 是一个功能强大、可扩展的报表生成工具，设计用于从多个数据源创建各种类型的报表。

## 概述

GenRep 是一个使用 Go 语言编写的命令行工具，基于 JSON 配置文件生成报表。它的设计特点包括：

1. **可扩展性**：通过插件架构可以轻松添加新的报表类型
2. **多功能性**：支持多种数据源和报表输出格式
3. **程序化**：可以作为实用工具被其他程序调用
4. **自包含**：所有配置都包含在单个 JSON 文件中

## 使用方法

1. 首先，确保你已经安装了 Go 1.20 或更高版本

1. 克隆项目仓库：

```bash
git clone https://github.com/genrep/ReportCraft.git
```

1. 进入项目目录并构建项目：

```bash
# macOS/Linux
cd ReportCraft
go build -o reportcraft ./cmd/reportcraft/

# Windows
cd ReportCraft
go build -o reportcraft.exe ./cmd/reportcraft/
```

1. 运行程序：

```bash
# macOS/Linux
./reportcraft -config=path/to/config.json

# Windows
reportcraft.exe -config=path\to\config.json
```

1. 生成的报告将保存在配置文件中指定的输出路径

## 配置格式

配置文件是一个具有以下结构的 JSON 文件：

```json
{
  "reportType": "string",           // 报表类型
  "outputPath": "string",           // 报表保存路径
  "outputFormat": "string",         // 输出格式 (docx, xlsx 等)
  "parameters": {                   // 报表特定参数
    // 根据报表类型有所不同
  },
  "dataSources": [                  // 数据源数组
    {
      "id": "string",               // 数据源唯一标识符
      "type": "string",             // 数据源类型 (db, api, file, influxdb, minio 等)
      "config": {                   // 数据源特定配置
        // 根据数据源类型有所不同
      }
    }
  ]
}
```

## 表格报告功能

ReportCraft 现在支持从不同来源生成数据表格：

1. 外部 JSON 文件：可以从 JSON 文件中检索特定路径的数据
2. 外部 CSV 文件：可以直接从 CSV 文件导入数据
3. 嵌入数据：可以直接在配置文件中包含数据

### 示例配置

`examples` 目录中提供了示例配置文件：

- `table-report-json.json`：展示如何使用外部 JSON 文件中的数据
- `table-report-csv.json`：展示如何使用外部 CSV 文件中的数据
- `table-report-embedded.json`：展示如何在配置中直接嵌入数据

### 表格配置选项

表格报告允许以下配置选项：

```json
{
  "tableConfig": {
    "title": "表格标题",
    "columns": [
      {"field": "列名", "header": "列标题", "width": 15},
      {"field": "另一列", "header": "另一个标题", "width": 20, "format": "date"}
    ],
    "headerStyle": {
      "bold": true,
      "background": "#DDEBF7",
      "color": "#000000"
    },
    "alternateRowStyle": true
  }
}
```

## 波形频谱报告

波形频谱报告功能允许生成包含波形和频谱数据的分析报告，适用于振动分析、声学分析等场景。

## 振动分析功能

ReportCraft 支持生成设备振动波形和频谱分析报告，包含以下功能：

1. 设备信息展示：设备名称、设备ID、测点位置等
2. 波形数据图表：显示时域振动波形
3. 频谱数据图表：显示频域分析结果
4. 特征参数表格：展示波形和频谱的特征参数

### 示例配置

`examples` 目录中提供了示例配置文件：

- `spectrum-report-embedded.json`：展示如何创建振动分析报告，使用内嵌数据
- `spectrum-report-csv.json`：展示如何创建振动分析报告，使用外部CSV文件数据

### 配置选项

振动分析报告支持以下配置选项：

```json
{
  "deviceInfo": {
    "deviceName": "设备名称",
    "deviceId": "设备ID",
    "location": "位置信息",
    "measurementPoint": "测点信息"
  },
  "waveformConfig": {
    "title": "波形图标题",
    "xLabel": "X轴标签",
    "yLabel": "Y轴标签"
  },
  "spectrumConfig": {
    "title": "频谱图标题",
    "xLabel": "X轴标签",
    "yLabel": "Y轴标签"
  }
}
```

## Windows环境注意事项

在Windows环境中使用ReportCraft时，需要注意以下几点：

1. 文件路径分隔符：Windows使用反斜杠(`\`)作为路径分隔符，而配置文件中建议使用正斜杠(`/`)，两者都可以正常工作。

2. 输出文件命名：建议使用英文文件名，避免使用中文或特殊字符作为文件名，以防止编码问题。这就是为什么示例配置文件中的输出路径使用了英文名称。

3. 命令行运行：在Windows环境中，使用`reportcraft.exe`而不是`./reportcraft`来运行程序。

4. 配置文件示例：所有示例配置文件中的路径格式都兼容Windows环境。

## 架构

GenRep 采用插件架构设计，允许轻松扩展新的报表类型和数据源：

1. **核心引擎**：处理配置解析、插件管理和协调
2. **报表插件**：实现特定的报表生成逻辑
3. **数据源适配器**：连接并从各种来源检索数据
4. **输出格式化器**：将报表数据转换为所需的输出格式

## 添加新的报表类型

通过实现 `ReportGenerator` 接口并将实现注册到插件系统，可以添加新的报表类型。

## 支持的数据源

- 数据库 (SQL, NoSQL)
- API (REST, GraphQL)
- 文件 (CSV, JSON, Excel)
- InfluxDB 2.0 时序数据库
- MinIO 对象存储
- 通过插件系统自定义数据源

## 支持的输出格式

- Microsoft Word (DOCX)
- Microsoft Excel (XLSX)
- PDF
- HTML
- 通过插件系统自定义格式
