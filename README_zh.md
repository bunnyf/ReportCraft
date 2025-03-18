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

2. 克隆项目仓库：

```bash
git clone https://github.com/genrep/ReportCraft.git
```

3. 进入项目目录并构建项目：

```bash
cd ReportCraft
go build -o reportcraft ./cmd/reportcraft/
```

4. 运行程序：

```bash
./reportcraft -config=path/to/config.json
```

5. 生成的报告将保存在配置文件中指定的输出路径

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
