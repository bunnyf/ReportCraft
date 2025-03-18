# GenRep 用户指南

## 简介

GenRep 是一个强大、可扩展的报表生成工具，设计用于从多种数据源创建各种类型的报表。它通过命令行界面操作，并支持通过插件进行扩展。

## 安装

1. 确保已安装 Go 1.20 或更高版本
2. 克隆仓库: `git clone https://github.com/genrep/genrep.git`
3. 构建应用程序: `cd genrep && go build -o bin/genrep cmd/genrep/main.go`

## 基本用法

GenRep 需要一个 JSON 配置文件来生成报表：

```
./bin/genrep --config path/to/config.json
```

## 配置文件格式

配置文件采用 JSON 格式，包含以下字段：

```json
{
  "report_type": "报表类型名称",
  "output_path": "输出文件路径",
  "output_format": "输出格式 (xlsx, docx 等)",
  "parameters": {
    "参数1": "值1",
    "参数2": "值2"
  },
  "data_sources": [
    {
      "name": "数据源名称",
      "type": "数据源类型",
      "config": {
        "配置参数1": "值1",
        "配置参数2": "值2"
      }
    }
  ]
}
```

## 支持的数据源

GenRep 支持以下数据源：

1. **文件数据源**
   - 类型: `file`
   - 配置参数:
     - `path`: 文件路径
     - `format`: 文件格式 (json, csv, etc.)

2. **InfluxDB 2.0 数据源**
   - 类型: `influxdb`
   - 配置参数:
     - `url`: InfluxDB 服务器 URL
     - `token`: 认证令牌
     - `org`: 组织名称

3. **MinIO 数据源**
   - 类型: `minio`
   - 配置参数:
     - `endpoint`: MinIO 服务器端点
     - `access_key`: 访问密钥
     - `secret_key`: 秘密密钥
     - `use_ssl`: 是否使用 SSL (true/false)
     - `bucket`: 存储桶名称

## 支持的报表类型

1. **Hello Word 报表**
   - 类型: `hello-word`
   - 描述: 生成一个包含 "Hello World" 文本的简单 Word 文档

2. **Hello Excel 报表**
   - 类型: `hello-excel`
   - 描述: 生成一个在 Sheet1 中包含 "Hello World" 文本的简单 Excel 文件

3. **波形报表**
   - 类型: `waveform`
   - 描述: 生成包含波形数据可视化的报表

## 示例

### 生成 Hello Word 报表

配置文件 `examples/hello-word.json`:

```json
{
  "report_type": "hello-word",
  "output_path": "./output/hello-world.docx",
  "output_format": "docx",
  "parameters": {},
  "data_sources": []
}
```

命令:

```
./bin/genrep --config examples/hello-word.json
```

### 生成 Hello Excel 报表

配置文件 `examples/hello-excel.json`:

```json
{
  "report_type": "hello-excel",
  "output_path": "./output/hello-world.xlsx",
  "output_format": "xlsx",
  "parameters": {},
  "data_sources": []
}
```

命令:

```
./bin/genrep --config examples/hello-excel.json
```

## 扩展 GenRep

GenRep 设计为可扩展的，可以通过实现适当的接口来添加新的数据源和报表类型。

### 添加新的数据源

1. 实现 `DataSource` 接口 (见 `internal/core/interfaces.go`)
2. 在 `main.go` 中注册新的数据源

### 添加新的报表类型

1. 实现 `ReportGenerator` 接口 (见 `internal/core/interfaces.go`)
2. 在 `main.go` 中注册新的报表生成器

## 故障排除

- 确保配置文件格式正确
- 检查数据源连接参数
- 验证输出目录存在并且可写
- 查看日志输出以获取详细的错误信息
