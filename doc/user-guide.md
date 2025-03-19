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

4. **频谱报表**
   - 类型: `SpectrumReport`
   - 描述: 生成包含频域数据可视化和设备振动分析的报表
   - 支持多种图表类型: 线图、柱状图、热图、3D表面图、瀑布图

## 支持的图表类型

ReportCraft支持以下图表类型：

### 基础图表类型

1. **线图**
   - 类型: `line`
   - 适用场景: 时域波形、趋势分析
   - 配置示例:

     ```json
     "chartType": "line",
     "chartStyle": {
       "lineColor": "#1E90FF",
       "lineWidth": 1.5,
       "markers": {
         "show": true,
         "size": 4
       }
     }
     ```

2. **柱状图**
   - 类型: `bar`
   - 适用场景: 频谱分析
   - 配置示例:

     ```json
     "chartType": "bar",
     "chartStyle": {
       "lineColor": "#32CD32",
       "lineWidth": 1.0
     }
     ```

3. **散点图**
   - 类型: `scatter`
   - 适用场景: 相关性分析
   - 配置示例:

     ```json
     "chartType": "scatter",
     "chartStyle": {
       "markerColor": "#FF4500",
       "markerSize": 5
     }
     ```

### 高级图表类型

1. **热图**
   - 类型: `heatmap`
   - 描述: 用于显示二维空间内的强度变化，如频率-时间-幅值的关系
   - 配置示例:

     ```json
     "chartType": "heatmap",
     "chartStyle": {
       "heatmap": {
         "colorScale": "viridis",  // 可选值: viridis, jet, plasma, inferno
         "showColorBar": true,     // 是否显示色彩条
         "interpolate": true       // 是否插值平滑
       }
     }
     ```

   - 数据源要求: 需要提供x(频率)、y(时间)、z(幅值)三维数据

2. **3D表面图**
   - 类型: `3dsurface`
   - 描述: 提供完整的三维数据可视化，展示频率-时间-幅值的立体关系
   - 配置示例:

     ```json
     "chartType": "3dsurface",
     "chartStyle": {
       "3dsurface": {
         "wireframe": true,          // 是否显示线框网格
         "colorScale": "jet",        // 色彩映射方案：jet, viridis, plasma, inferno, rainbow
         "colorOpacity": 0.85,       // 曲面颜色透明度(0-1)
         "meshDensity": {            // 网格密度设置
           "x": 50,                  // X轴网格密度
           "y": 50                   // Y轴网格密度
         },
         "rotation": {               // 初始视角设置
           "x": 30,                  // X轴旋转角度
           "y": 45,                  // Y轴旋转角度
           "z": 0                    // Z轴旋转角度
         },
         "lighting": {               // 光照效果
           "ambient": 0.6,           // 环境光强度
           "diffuse": 0.8,           // 漫反射强度
           "specular": 0.2           // 镜面反射强度
         },
         "interactiveControls": true // 是否启用交互式控制(旋转、缩放)
       }
     }
     ```

   - 数据源要求: 需要提供x(频率)、y(时间)、z(幅值)三维数据
   - 最佳实践:
     - 针对大型数据集，可以调整`meshDensity`降低渲染复杂度提高性能
     - 使用`colorScale`和`lighting`配置强调重要数据区域
     - 启用`interactiveControls`允许用户自由旋转和缩放图表，便于深入分析

3. **瀑布图**
   - 类型: `waterfall`
   - 描述: 用于分析随时间或其他条件变化的频谱，展示设备振动状态的演变
   - 配置示例:

     ```json
     "chartType": "waterfall",
     "chartStyle": {
       "waterfall": {
         "baseColor": "#1E90FF",      // 基础颜色
         "colorGradient": true,       // 是否使用渐变色
         "gradientScheme": "rainbow", // 渐变色方案：rainbow, thermal, terrain, blues
         "spacing": 0.1,              // 图层间距(0-1)
         "perspective": 30,           // 透视角度(0-60)
         "lineWidth": 1.2,            // 线条宽度
         "fillOpacity": 0.7,          // 填充透明度(0-1)
         "showGridLines": true,       // 显示网格线
         "gridColor": "#333333",      // 网格线颜色
         "labels": {                  // 标签设置
           "showZValues": true,       // 在图层上显示Z值标签
           "fontColor": "#EEEEEE",    // 标签文字颜色
           "fontSize": 10             // 标签文字大小
         },
         "yAxisMode": "time",         // Y轴模式：time(时间)或order(阶次)
         "renderQuality": "high"      // 渲染质量：low, medium, high
       }
     }
     ```

   - 数据源要求: 需要提供x(频率)、y(时间)、z(幅值)三维数据
   - 应用场景:
     - 启动/关闭分析：捕捉设备启动或关闭过程中的振动变化
     - 转速变化监测：分析不同转速下的振动频谱变化
     - 故障发展追踪：跟踪故障从发生到发展的过程中频谱特征的变化

## 图表通用配置项

所有图表类型都支持以下通用配置选项：

```json
"chartStyle": {
  "axis": {
    "xMin": 0,            // X轴最小值
    "xMax": 1000,         // X轴最大值
    "yMin": 0,            // Y轴最小值
    "yMax": 0.1,          // Y轴最大值
    "zMin": 0,            // Z轴最小值(仅3D图表)
    "zMax": 0.1           // Z轴最大值(仅3D图表)
  },
  "grid": {
    "show": true,         // 是否显示网格
    "color": "#CCCCCC",   // 网格颜色
    "lineStyle": "solid", // 线型: solid, dashed, dotted
    "lineWidth": 0.5      // 线宽
  },
  "labels": {
    "title": {
      "fontSize": 14,      // 标题字体大小
      "fontWeight": "bold",// 字体粗细
      "color": "#333333"   // 文字颜色
    },
    "axis": {
      "fontSize": 12,      // 轴标签字体大小
      "color": "#666666"   // 轴标签颜色
    }
  }
}
```

## 示例

### 生成 Hello Word 报表

配置文件 `examples/hello-word.json`：

```json
{
  "report_type": "hello-word",
  "output_path": "./output/hello-world.docx",
  "output_format": "docx",
  "parameters": {},
  "data_sources": []
}
```

命令：

```bash
./bin/genrep --config examples/hello-word.json
```

### 生成 Hello Excel 报表

配置文件 `examples/hello-excel.json`：

```json
{
  "report_type": "hello-excel",
  "output_path": "./output/hello-world.xlsx",
  "output_format": "xlsx",
  "parameters": {},
  "data_sources": []
}
```

命令：

```bash
./bin/genrep --config examples/hello-excel.json
```

### 生成频谱分析报表（使用CSV数据源）

配置文件 `examples/spectrum-report-csv.json`：

```json
{
  "report_type": "SpectrumReport",
  "output_path": "./output/spectrum-report-csv.xlsx",
  "output_format": "xlsx",
  "parameters": {},
  "data_sources": [
    {
      "name": "csv_data",
      "type": "file",
      "config": {
        "path": "./data/spectrum-data.csv",
        "format": "csv"
      }
    }
  ]
}
```

命令：

```bash
./bin/genrep --config examples/spectrum-report-csv.json
```

这将生成一份包含基本波形和频谱图表，以及热图、3D表面图和瀑布图的完整报表。

### 生成频谱分析报表（使用内嵌数据）

配置文件 `examples/spectrum-report-embedded.json`：

```json
{
  "report_type": "SpectrumReport",
  "output_path": "./output/spectrum-report-embedded.xlsx",
  "output_format": "xlsx",
  "parameters": {},
  "data_sources": [
    {
      "name": "embedded_data",
      "type": "embedded",
      "config": {
        "data": [
          [1, 2, 3],
          [4, 5, 6],
          [7, 8, 9]
        ]
      }
    }
  ]
}
```

命令：

```bash
./bin/genrep --config examples/spectrum-report-embedded.json
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
