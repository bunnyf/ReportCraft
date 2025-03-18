# GenRep 开发者指南

本文档提供了关于 GenRep 代码结构、架构和开发工作流程的详细信息，旨在帮助新的贡献者理解项目并进行扩展。

## 项目架构

GenRep 使用模块化、插件式的架构，可以轻松扩展新的数据源、报表类型和输出格式。

```
GenRep/
├── cmd/              # 命令行入口
├── doc/              # 文档
├── examples/         # 示例配置文件
├── internal/         # 内部实现
│   ├── core/         # 核心接口和引擎
│   ├── datasource/   # 数据源实现
│   ├── formatter/    # 输出格式实现
│   ├── logger/       # 日志系统
│   ├── plugin/       # 插件系统
│   └── report/       # 报表生成器实现
├── output/           # 报表输出目录
└── tests/            # 集成测试
```

## 核心组件

### 1. 报表引擎 (`internal/core/engine.go`)

报表引擎是 GenRep 的核心，负责协调数据源、报表生成器和输出格式化器。它提供以下功能：

- 注册和管理数据源
- 注册和管理报表生成器
- 注册和管理输出格式化器
- 解析配置文件
- 协调报表生成流程

### 2. 接口定义 (`internal/core/interfaces.go`)

该文件定义了系统中的主要接口：

- `DataSource`: 数据源接口
- `ReportGenerator`: 报表生成器接口
- `OutputFormatter`: 输出格式化器接口

### 3. 数据源 (`internal/datasource/`)

数据源负责从各种来源获取数据。所有数据源必须实现 `DataSource` 接口。当前支持的数据源：

- `FileDataSource`: 从文件读取数据
- `InfluxDBDataSource`: 从 InfluxDB 2.0 获取数据
- `MinioDataSource`: 从 MinIO 获取对象

### 4. 报表生成器 (`internal/report/`)

报表生成器负责生成特定类型的报表。所有报表生成器必须实现 `ReportGenerator` 接口。当前支持的报表：

- `HelloWordReport`: 生成简单的 Word 文档
- `HelloExcelReport`: 生成简单的 Excel 文档
- `WaveformReport`: 生成波形数据报表

### 5. 输出格式化器 (`internal/formatter/`)

输出格式化器负责将数据转换为特定的输出格式。当前支持的格式：

- `DocxFormatter`: Word 文档格式
- `ExcelFormatter`: Excel 文档格式

## 开发工作流程

### 设置开发环境

1. 克隆仓库：
   ```bash
   git clone https://github.com/genrep/genrep.git
   cd genrep
   ```

2. 安装依赖：
   ```bash
   go mod tidy
   ```

3. 构建应用程序：
   ```bash
   go build -o bin/genrep cmd/genrep/main.go
   ```

### 添加新的数据源

1. 在 `internal/datasource/` 创建新文件，例如 `mysql.go`
2. 实现 `DataSource` 接口
3. 提供一个工厂函数，如 `NewMySQLDataSource()`
4. 在 `cmd/genrep/main.go` 中注册数据源：
   ```go
   engine.RegisterDataSource("mysql", datasource.NewMySQLDataSource)
   ```

示例实现:

```go
package datasource

import (
    "context"
    "database/sql"
    "fmt"
    
    _ "github.com/go-sql-driver/mysql"
)

type MySQLDataSource struct {
    db *sql.DB
}

func (ds *MySQLDataSource) Initialize(config map[string]interface{}) error {
    // 提取连接参数
    dsn, ok := config["dsn"].(string)
    if !ok {
        return fmt.Errorf("dsn is required")
    }
    
    // 建立连接
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return fmt.Errorf("failed to connect to MySQL: %w", err)
    }
    
    ds.db = db
    return nil
}

func (ds *MySQLDataSource) FetchData(ctx context.Context, query interface{}) (interface{}, error) {
    // 实现查询逻辑
    // ...
}

func (ds *MySQLDataSource) Close() error {
    if ds.db != nil {
        return ds.db.Close()
    }
    return nil
}

func NewMySQLDataSource() *MySQLDataSource {
    return &MySQLDataSource{}
}
```

### 添加新的报表类型

1. 在 `internal/report/` 创建新文件，例如 `pdf_report.go`
2. 实现 `ReportGenerator` 接口
3. 提供一个工厂函数，如 `NewPDFReport()`
4. 在 `cmd/genrep/main.go` 中注册报表生成器：
   ```go
   engine.RegisterReportGenerator("pdf-report", report.NewPDFReport)
   ```

### 测试

1. 单元测试：
   ```bash
   go test ./internal/...
   ```

2. 运行特定测试：
   ```bash
   go test ./internal/datasource -run TestMySQLDataSource
   ```

3. 测试覆盖率：
   ```bash
   go test ./internal/... -cover
   ```

## 代码规范

- 所有代码应遵循 Go 标准代码规范（使用 `gofmt`）
- 所有导出的函数、类型和变量应当有文档注释
- 使用 `go vet` 和 `golint` 检查代码质量
- 实现适当的错误处理和日志记录
- 编写单元测试，争取达到合理的测试覆盖率

## 提交变更

1. 创建分支：
   ```bash
   git checkout -b feature/my-new-feature
   ```

2. 提交变更：
   ```bash
   git commit -m "feat: add my new feature"
   ```

3. 推送分支：
   ```bash
   git push origin feature/my-new-feature
   ```

4. 创建合并请求

## 版本发布流程

1. 更新版本号（遵循语义化版本控制）
2. 更新 CHANGELOG.md
3. 创建版本标签
4. 构建和发布二进制文件

## 常见问题

### 问题：无法找到依赖项

解决方案：运行 `go mod tidy` 下载所有依赖项。

### 问题：报表生成失败

解决方案：检查配置文件、确保数据源配置正确、查看详细日志。

### 问题：构建错误

解决方案：确保 Go 版本正确（1.20 或更高）、检查代码错误。
