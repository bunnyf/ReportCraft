# GenRep - General Reporting Tool

GenRep is a powerful, extensible report generation tool designed to create various types of reports from multiple data sources.

## Overview

GenRep is a command-line tool written in Go that generates reports based on a JSON configuration file. It is designed to be:

1. **Extensible**: New report types can be easily added through a plugin architecture
2. **Versatile**: Supports multiple data sources and report output formats
3. **Programmatic**: Can be called by other programs as a utility
4. **Self-contained**: All configuration is contained in a single JSON file

## Usage

1. First, make sure you have Go 1.20 or higher installed

1. Clone the repository:

```bash
git clone https://github.com/genrep/ReportCraft.git
```

1. Navigate to the project directory and build the project:

```bash
# macOS/Linux
cd ReportCraft
go build -o reportcraft ./cmd/reportcraft/

# Windows
cd ReportCraft
go build -o reportcraft.exe ./cmd/reportcraft/
```

1. Run the program:

```bash
# macOS/Linux
./reportcraft -config=path/to/config.json

# Windows
reportcraft.exe -config=path\to\config.json
```

1. The generated report will be saved at the output path specified in the config file

## Configuration Format

The configuration file is a JSON file with the following structure:

```json
{
  "reportType": "string",           // Type of report to generate
  "outputPath": "string",           // Where to save the generated report
  "outputFormat": "string",         // Output format (docx, xlsx, etc.)
  "parameters": {                   // Report-specific parameters
    // Varies based on report type
  },
  "dataSources": [                  // Array of data sources
    {
      "id": "string",               // Unique identifier for this data source
      "type": "string",             // Type of data source (db, api, file, etc.)
      "config": {                   // Data source-specific configuration
        // Varies based on data source type
      }
    }
  ]
}
```

## Table Report Feature

ReportCraft now supports generating data tables from different sources:

1. External JSON file: Data can be retrieved from a JSON file with a specific data path
2. External CSV file: Data can be imported directly from a CSV file
3. Embedded data: Data can be included directly in the configuration file

### Example Configurations

Example configuration files are available in the `examples` directory:

- `table-report-json.json`: Shows how to use data from an external JSON file
- `table-report-csv.json`: Shows how to use data from an external CSV file
- `table-report-embedded.json`: Shows how to embed data directly in the configuration

### Table Configuration Options

The table report allows for the following configuration options:

```json
{
  "tableConfig": {
    "title": "Table Title",
    "columns": [
      {"field": "columnName", "header": "Column Header", "width": 15},
      {"field": "anotherColumn", "header": "Another Header", "width": 20, "format": "date"}
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

## 振动分析报告

ReportCraft supports generating vibration waveform and spectrum analysis reports, with the following features:

1. Device information display: device name, device ID, measurement point, etc.
2. Waveform data chart: displaying time domain vibration waveform
3. Spectrum data chart: displaying frequency domain spectrum data
4. Parameter tables: displaying waveform and spectrum characteristic parameters

### 图表样式配置

The vibration analysis report offers rich chart style configuration options:

#### 基础样式配置

```json
"chartStyle": {
  "lineColor": "#1E90FF",
  "lineWidth": 1.5,
  "gridLines": true
}
```

#### 增强样式配置

```json
"chartStyle": {
  // 基础样式
  "lineColor": "#1E90FF",
  "lineWidth": 1.5,
  "gridLines": true,
  
  // 坐标轴范围
  "axis": {
    "xMin": 0,
    "xMax": 100,
    "yMin": -0.5,
    "yMax": 0.5
  },
  
  // 数据点标记
  "markers": {
    "show": true,
    "size": 4,
    "color": "#FF4500",
    "shape": "circle"  // 可选: circle, square, triangle, diamond
  },
  
  // 网格线配置
  "grid": {
    "show": true,
    "color": "#CCCCCC",
    "lineStyle": "solid", // 可选: solid, dashed, dotted
    "lineWidth": 0.5,
    "minorGrid": {
      "show": true,
      "color": "#EEEEEE",
      "lineStyle": "dotted",
      "lineWidth": 0.25
    }
  },
  
  // 突出显示区域
  "highlight": {
    "regions": [
      {
        "xStart": 40,
        "xEnd": 60,
        "color": "rgba(255, 100, 100, 0.2)",
        "label": "特征区域"
      }
    ]
  },
  
  // 标题和标签样式
  "labels": {
    "title": {
      "fontSize": 14,
      "fontWeight": "bold",
      "color": "#333333"
    },
    "axis": {
      "fontSize": 12,
      "color": "#666666"
    }
  },
  
  // 图例配置
  "legend": {
    "position": "bottom", // 可选: top, bottom, left, right
    "fontSize": 11,
    "color": "#333333"
  }
}
```

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

2. 输出文件命名：建议使用英文文件名，避免使用中文或特殊字符作为文件名，以防止编码问题。

3. 命令行运行：在Windows环境中，使用`reportcraft.exe`而不是`./reportcraft`来运行程序。

4. 配置文件示例：所有示例配置文件中的路径格式都兼容Windows环境。

## Architecture

GenRep is designed with a plugin architecture that allows for easy extension with new report types and data sources:

1. **Core Engine**: Handles configuration parsing, plugin management, and orchestration
2. **Report Plugins**: Implement specific report generation logic
3. **Data Source Adapters**: Connect to and retrieve data from various sources
4. **Output Formatters**: Convert report data to the desired output format

## Adding New Report Types

New report types can be added by implementing the `ReportGenerator` interface and registering the implementation with the plugin system.

## Supported Data Sources

- Databases (SQL, NoSQL)
- APIs (REST, GraphQL)
- Files (CSV, JSON, Excel)
- Custom data sources through the plugin system

## Supported Output Formats

- Microsoft Word (DOCX)
- Microsoft Excel (XLSX)
- PDF
- HTML
- Custom formats through the plugin system
